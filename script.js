// Fungsi untuk menghitung jarak menggunakan rumus Haversine
function haversineDistance(coords1, coords2) {
    function toRad(x) {
        return x * Math.PI / 180;
    }

    var lat1 = coords1[0];
    var lon1 = coords1[1];
    var lat2 = coords2[0];
    var lon2 = coords2[1];

    var R = 6371; // Radius bumi dalam kilometer
    var x1 = lat2 - lat1;
    var dLat = toRad(x1);
    var x2 = lon2 - lon1;
    var dLon = toRad(x2)
    var a = Math.sin(dLat / 2) * Math.sin(dLat / 2) +
        Math.cos(toRad(lat1)) * Math.cos(toRad(lat2)) *
        Math.sin(dLon / 2) * Math.sin(dLon / 2);
    var c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    var d = R * c;

    return d; // Jarak dalam kilometer
}

// Fungsi untuk memuat data McDonald's dari API backend
function loadMcDonalds(userCoords) {
    fetch('http://localhost:8080/mcdonalds')
        .then(response => response.json())
        .then(data => {
            data.data.mcdonalds.forEach(mcd => {
                var mcdCoords = [mcd.latitude, mcd.longitude];
                var distance = haversineDistance(userCoords, mcdCoords).toFixed(2); // Hasil dalam km

                L.marker(mcdCoords)
                    .addTo(map)
                    .bindPopup(`<b>${mcd.name}</b><br>Distance: ${distance} km`);
            });
        })
        .catch(error => console.error('Error loading McDonalds:', error));
}

// Fungsi untuk menemukan lokasi pengguna
function locateUser() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(position => {
            var latitude = position.coords.latitude;
            var longitude = position.coords.longitude;
            var userCoords = [latitude, longitude];

            map.setView(userCoords, 15);
            L.marker(userCoords)
                .addTo(map)
                .bindPopup("You are here")
                .openPopup();

            // Memuat data McDonald's setelah lokasi pengguna ditemukan
            loadMcDonalds(userCoords);
        }, () => {
            alert("Unable to retrieve your location.");
        });
    } else {
        alert("Geolocation is not supported by your browser.");
    }
}

// Inisialisasi peta
var map = L.map('map').setView([0, 0], 2);

// Menambahkan tile layer dari OpenStreetMap
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

// Menjalankan locateUser saat halaman dimuat
locateUser();
