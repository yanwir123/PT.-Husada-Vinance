create database HusadaVinance;

CREATE TABLE Keuangan (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Hari VARCHAR(255),
    Transaksi_Masuk DECIMAL(15,2),
    Transaksi_Keluar DECIMAL(15,2),
    Total DECIMAL(15,2),
    Bulan VARCHAR(255),
    Jumlah_Transaksi DECIMAL(15,2),
    Total_Transaksi DECIMAL(15,2),
    Keterangan VARCHAR(255)
);
