const apiUrl = 'http://localhost:8080/mcdonalds'; // Replace with your actual API URL
let map;
let userMarker;
let storeMarkers = [];

function initMap() {
    map = L.map('map').setView([0, 0], 13); // Center map at [0,0] initially

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);
}

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(findNearestMcDonald, showError);
    } else {
        document.getElementById('result').innerHTML = "Geolocation is not supported by this browser.";
    }
}

function findNearestMcDonald(position) {
    const userLat = position.coords.latitude;
    const userLon = position.coords.longitude;

    // Center the map on the user's location
    if (userMarker) {
        map.removeLayer(userMarker);
    }
    userMarker = L.marker([userLat, userLon]).addTo(map);
    map.setView([userLat, userLon], 13);

    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            const stores = data.data.mcdonalds;
            if (stores.length === 0) {
                document.getElementById('result').innerHTML = "No McDonald's locations found.";
                return;
            }

            let nearestStore = null;
            let shortestDistance = Infinity;

            // Clear previous store markers
            storeMarkers.forEach(marker => map.removeLayer(marker));
            storeMarkers = [];

            stores.forEach(store => {
                const distance = calculateDistance(userLat, userLon, store.latitude, store.longitude);
                if (distance < shortestDistance) {
                    shortestDistance = distance;
                    nearestStore = store;
                }

                // Add store marker to the map
                const marker = L.marker([store.latitude, store.longitude])
                    .bindPopup(`<b>${store.name}</b><br>Distance: ${distance.toFixed(2)} km`)
                    .addTo(map);
                storeMarkers.push(marker);
            });

            if (nearestStore) {
                document.getElementById('result').innerHTML = `
                    Nearest McDonald's:
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
            console.error('Error fetching McDonald\'s data:', error);
            document.getElementById('result').innerHTML = "Error fetching McDonald's data.";
        });
}

function calculateDistance(lat1, lon1, lat2, lon2) {
    const R = 6371; // Radius of the Earth in km
    const dLat = toRadians(lat2 - lat1);
    const dLon = toRadians(lon2 - lon1);
    const a = Math.sin(dLat / 2) * Math.sin(dLat / 2) +
              Math.cos(toRadians(lat1)) * Math.cos(toRadians(lat2)) *
              Math.sin(dLon / 2) * Math.sin(dLon / 2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    return R * c; // Distance in km
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

// Initialize the map when the page loads
window.onload = initMap;
