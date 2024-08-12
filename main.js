const apiUrl = 'http://localhost:8080/stores'; // Ganti dengan URL API Anda
let map;
let userMarker;
let storeMarkers = [];

function initMap() {
    map = L.map('map').setView([0, 0], 13); // Pusatkan peta pada [0,0] awalnya

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);
}

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(findNearestStore, showError);
    } else {
        document.getElementById('result').innerHTML = "Geolocation is not supported by this browser.";
    }
}

function findNearestStore(position) {
    const userLat = position.coords.latitude;
    const userLon = position.coords.longitude;

    // Pusatkan peta pada lokasi pengguna
    if (userMarker) {
        map.removeLayer(userMarker);
    }
    userMarker = L.marker([userLat, userLon]).addTo(map);
    map.setView([userLat, userLon], 13);

    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            const stores = data.stores; // Asumsikan data berbentuk { stores: [{name, latitude, longitude}, ...]}
            if (stores.length === 0) {
                document.getElementById('result').innerHTML = "No stores found.";
                return;
            }

            let nearestStore = null;
            let shortestDistance = Infinity;

            // Hapus marker toko sebelumnya
            storeMarkers.forEach(marker => map.removeLayer(marker));
            storeMarkers = [];

            stores.forEach(store => {
                const distance = calculateDistance(userLat, userLon, store.latitude, store.longitude);
                if (distance < shortestDistance) {
                    shortestDistance = distance;
                    nearestStore = store;
                }

                // Tambahkan marker toko ke peta
                const marker = L.marker([store.latitude, store.longitude])
                    .bindPopup(`<b>${store.name}</b><br>Distance: ${distance.toFixed(2)} km`)
                    .addTo(map);
                storeMarkers.push(marker);
            });

            if (nearestStore) {
                document.getElementById('result').innerHTML = `
                    Nearest Store:
                    <ul>
                        <li>Name: ${nearestStore.name}</li>
                        <li>Latitude: ${nearestStore.latitude}</li>
                        <li>Longitude: ${nearestStore.longitude}</li>
                        <li>Distance: ${shortestDistance.toFixed(2)} km</li>
                    </ul>
                `;
            }
        })
        .catch(error => {
            console.error('Error fetching store data:', error);
            document.getElementById('result').innerHTML = "Error fetching store data.";
        });
}

function calculateDistance(lat1, lon1, lat2, lon2) {
    const R = 6371; // Radius bumi dalam km
    const dLat = toRadians(lat2 - lat1);
    const dLon = toRadians(lon2 - lon1);
    const a = Math.sin(dLat / 2) * Math.sin(dLat / 2) +
              Math.cos(toRadians(lat1)) * Math.cos(toRadians(lat2)) *
              Math.sin(dLon / 2) * Math.sin(dLon / 2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    return R * c; // Jarak dalam km
}

function toRadians(degrees) {
    return degrees * (Math.PI / 180);
}

function showError(error) {
    let errorMessage = "An unknown error occurred.";
    switch(error.code) {
        case error.PERMISSION_DENIED:
            errorMessage = "User denied the request for Geolocation.";
            break;
        case error.POSITION_UNAVAILABLE:
            errorMessage = "Location information is unavailable.";
            break;
        case error.TIMEOUT:
            errorMessage = "The request to get user location timed out.";
            break;
        case error.UNKNOWN_ERROR:
            errorMessage = "An unknown error occurred.";
            break;
    }
    document.getElementById('result').innerHTML = errorMessage;
}

// Inisialisasi peta saat halaman dimuat
window.onload = initMap;
