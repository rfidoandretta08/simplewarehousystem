Tugas ini merupakan kelanjutan dari tugas sebelumnya yang berjudul **"Web Server dan Golang Route"**. Pada tugas ini, kami akan memperluas fitur backend yang telah dibuat sebelumnya dengan menambahkan file handling dan storage API untuk mendukung unggah dan unduh file gambar produk.

Tugas ini berfokus pada implementasi file handling dan storage API. Fitur yang dikembangkan dalam tugas ini adalah sebagai berikut:

- **Endpoint Upload**: Mengunggah gambar produk ke penyimpanan lokal.
- **Endpoint Download**: Mengunduh gambar produk berdasarkan ID produk.
- **Validasi File**: Memastikan hanya format gambar tertentu yang diterima (misal: PNG, JPG, JPEG) dengan ukuran maksimum yang ditentukan.
- **Error Handling**: Menangani skenario kesalahan seperti file berukuran terlalu besar atau format yang tidak sesuai.

## Fitur Utama

### 1. **File Handling and Storage**
- **Endpoint Upload**: Mengunggah gambar produk ke penyimpanan lokal.
  - Endpoint: `POST /products/upload/{id}`
  - contohnya : http://localhost:8080/products/upload/2 lalu isi body dengan format form-data
  - form data : key = file || value adalaha file yang anda ingin upload (example : kulkas.jpg)
  - Gambar yang diunggah akan disimpan dalam folder lokal `/uploads`.
  - Format gambar yang diterima: JPG dan JPEG.
  - Ukuran maksimal file: 1 MB.
  
- **Endpoint Download**: Mengunduh gambar produk berdasarkan ID produk.
  - Endpoint: `GET /products/file/download/{id}?URL gambar`
  - contoh : http://localhost:8080/products/file/download/2?file_name=kulkas.jpeg&directory_name=uploads
  - Gambar produk diunduh berdasarkan ID produk yang diminta.

### 2. **Web Server and Routing**
- Framework yang digunakan: **Gin** (Golang web framework).
- Menangani metode HTTP `POST` dan `GET` untuk endpoint unggah dan unduh gambar.
  
### 3. **API Integration and Testing**
- API diuji menggunakan **Postman** untuk memastikan fungsionalitas unggah dan unduh file berfungsi dengan baik.
- Validasi input dan error handling diterapkan dengan baik untuk menangani kesalahan file seperti format tidak sesuai atau ukuran file terlalu besar.

## Tools

- **MySQL**: Untuk penyimpanan data produk (termasuk informasi file gambar).
- **Golang**: Bahasa pemrograman untuk membangun server dan API.
- **Gin**: Framework untuk routing API.
- **Postman**: Untuk menguji endpoint API.
- **File Storage**: Penyimpanan gambar produk di folder lokal.
