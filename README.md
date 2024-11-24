# Elasticsearch Example Project

Repository ini memberikan contoh praktis dan panduan penggunaan Elasticsearch, mulai dari instalasi, konfigurasi, hingga menjalankan operasi dasar.

---

## 📂 Struktur Folder

```plaintext
ex-elastic-search/
├── go-app/           # Contoh implementasi Elasticsearch menggunakan Go
├── nodejs-app/       # Contoh implementasi Elasticsearch menggunakan Node.js
└── shared/           # Resource yang digunakan bersama
```

---

## 🚀 Cara Instalasi dan Menjalankan Elasticsearch

### Langkah 1: **Instal Elasticsearch**
1. Unduh Elasticsearch dari [situs resmi Elasticsearch](https://www.elastic.co/downloads/elasticsearch).
2. Ekstrak file yang diunduh ke lokasi yang diinginkan.

### Langkah 2: **Jalankan Elasticsearch**
#### Pilihan 1: Jalankan Manual
1. Masuk ke folder instalasi Elasticsearch.
2. Buka **Command Prompt (CMD)** atau terminal.
3. Jalankan perintah berikut:
   ```bash
   bin\elasticsearch.bat
   ```

#### Pilihan 2: Jalankan dengan Docker
1. Pastikan Docker sudah terinstal.
2. Jalankan perintah berikut di terminal:
   ```bash
   docker run -p 9200:9200 -e "discovery.type=single-node" elasticsearch:8.x.x
   ```

### Langkah 3: **Periksa Status Elasticsearch**
- Buka browser dan akses:
    - **HTTP**: [http://localhost:9200](http://localhost:9200)
    - **HTTPS** (jika diaktifkan): [https://localhost:9200](https://localhost:9200)

---

## 🔑 Mengelola Username dan Password

### **Detail Default**
- **Username default**: `elastic`
- **Password default**: Dapat ditemukan di **log Elasticsearch** saat pertama kali dijalankan.

### **Mengganti Password User**
1. Masuk ke folder instalasi Elasticsearch.
2. Buka **CMD** dan jalankan perintah berikut:
   ```bash
   bin\elasticsearch-reset-password -u [username]
   ```
3. Ganti `[username]` dengan nama pengguna. Contoh:
   ```bash
   bin\elasticsearch-reset-password -u elastic
   ```
4. **Pilihan Password**:
    - Jika ingin menggunakan password tertentu, masukkan password saat diminta.
    - Jika dikosongkan, Elasticsearch akan otomatis menghasilkan password baru dan menampilkannya di terminal.

---

## ⚙️ Konfigurasi Elasticsearch

File konfigurasi utama dapat ditemukan di:

```plaintext
config/elasticsearch.yml
```

### Konfigurasi Utama:
- **Izinkan akses dari semua jaringan**:
  ```yaml
  http.host: 0.0.0.0
  ```

- **Ganti port HTTP default**:
  ```yaml
  http.port: 9200
  ```

- **Nonaktifkan HTTPS (hanya untuk pengembangan lokal)**:
  ```yaml
  xpack.security.http.ssl.enabled: false
  ```

- **Setel lokasi data dan log**:
  ```yaml
  path.data: /path/to/data
  path.logs: /path/to/logs
  ```

Setelah mengubah konfigurasi, **restart Elasticsearch** agar perubahan diterapkan.

---

## 📖 Dokumentasi Resmi dan Referensi

Berikut adalah sumber referensi penting untuk Elasticsearch:

1. [Dokumentasi Elasticsearch](https://www.elastic.co/guide/index.html)
2. [Query DSL Reference](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html)
3. [Elasticsearch Node.js Client](https://www.elastic.co/guide/en/elasticsearch/client/javascript-api/current/index.html)
4. [Elasticsearch Go Client](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/index.html)

---

Dengan panduan ini, Anda dapat memulai eksplorasi Elasticsearch dengan mudah. Selamat mencoba! 🚀