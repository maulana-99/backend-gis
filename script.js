// Inisialisasi peta
var map = L.map('map').setView([0, 0], 2);

// Menambahkan tile layer dari OpenStreetMap
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

// Fungsi untuk memuat data McDonald's dari API backend
function loadMcDonalds() {
    fetch('http://localhost:8080/mcdonalds')
        .then(response => response.json())
        .then(data => {
            data.data.mcdonalds.forEach(mcd => {
                L.marker([mcd.latitude, mcd.longitude])
                    .addTo(map)
                    .bindPopup(`<b>${mcd.name}</b>`);
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
            map.setView([latitude, longitude], 15);
            L.marker([latitude, longitude])
                .addTo(map)
                .bindPopup("You are here")
                .openPopup();
        }, () => {
            alert("Unable to retrieve your location.");
        });
    } else {
        alert("Geolocation is not supported by your browser.");
    }
}

// Memuat data McDonald's saat halaman dimuat
loadMcDonalds();
