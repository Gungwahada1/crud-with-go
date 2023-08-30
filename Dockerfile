# Gunakan image base yang mengandung runtime Go
FROM golang:1.20.0

# Setel direktori kerja di dalam kontainer
WORKDIR /app

# Salin file-file yang diperlukan ke dalam kontainer
COPY . .

# Kompilasi proyek Go
RUN go build -o main .

# Tetapkan perintah default untuk menjalankan saat kontainer berjalan
CMD ["./main"]