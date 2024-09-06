# Proyek MVC - McDonald's Locator

## Deskripsi

Proyek ini adalah aplikasi berbasis web untuk mengelola lokasi McDonald's menggunakan arsitektur Model-View-Controller (MVC). Aplikasi ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data McDonald's dan juga melakukan login/logout.

## Struktur Proyek BE

- `config/`: Berisi konfigurasi aplikasi, termasuk pengaturan koneksi database.
- `routers/`: Berisi konfigurasi rute dan pengaturan router.
- `internal/`: Berisi logika bisnis dan pengontrol untuk aplikasi.
  - `album/`: Modul Crud tester.
  - `mcdonal/`: Modul untuk mengelola data McDonald's.
  - `users/`: Modul untuk mengelola login dan logout pengguna.
- `main.go`: Titik masuk aplikasi, menginisialisasi database, dan memulai server HTTP.

## Struktur Proyek FE

- `public/`: Berisi aset statis seperti gambar, ikon, dan file HTML.
- `src/`: Berisi kode sumber aplikasi frontend.
  - `components/`: Berisi komponen-komponen Vue yang digunakan dalam aplikasi.
  - `views/`: Berisi halaman-halaman utama aplikasi.
  - `router/`: Berisi konfigurasi rute untuk aplikasi.
  - `assets/`: Berisi aset-aset seperti gambar dan ikon.
  - `App.vue`: Komponen utama aplikasi.
  - `main.js`: Titik masuk aplikasi frontend.

## Fitur

- **McDonald's Locator:**
  - Menampilkan daftar McDonald's
  - Menambahkan McDonald's baru
  - Mengupdate koordinat McDonald's
  - Menghapus McDonald's
  - Menampilkan jarak McDonald's dengan user

- **User Authentication:**
  - Login pengguna
  - Logout pengguna

## Instalasi

1. **Clone Repository Backend**

   ```bash
   git clone https://github.com/maulana-99/backend-gis.git
   cd backend-gis
   ```

2. **Instalasi Backend**

   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```

3. **Clone Repository Frontend**

   ```bash
   git clone https://github.com/maulana-99/frontend-gis.git
   cd frontend-gis
   ```

4. **Instalasi Frontend**

   ```bash
   cd frontend
   npm install
   npm run dev
   ```
