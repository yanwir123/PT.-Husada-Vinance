package repository

import (
	models "PT.HusadaVinance/Models"
	connection "PT.HusadaVinance/Models/Connection"
	DataPerusahaan "PT.HusadaVinance/Models/DataPerusahaan"
)

// InsertJurusan melakukan penyisipan data ke tabel Keuangan.
func InsertHusadaVinance(DP DataPerusahaan.Keuangan) models.BaseResponseModels {
	var query string
	var result models.BaseResponseModels
	db := connection.DB

	// Query SQL untuk menyisipkan data ke tabel Keuangan
	query = "INSERT INTO Keuangan (Hari, Transaksi_Masuk, Transaksi_Keluar, Total, Bulan, Jumlah_Transaksi, Total_Transaksi, Keterangan) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	// Melakukan eksekusi query dengan data yang diberikan
	tempResult := db.Exec(query, DP.Hari, DP.Transaksi_Masuk, DP.Transaksi_Keluar, DP.Total, DP.Bulan, DP.Jumlah_Transaksi, DP.Total_Transaksi, DP.Keterangan)

	// Memeriksa apakah terdapat kesalahan selama eksekusi query
	if tempResult.Error != nil {
		result = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Jika eksekusi query berhasil tanpa kesalahan
		result = models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data berhasil ditambahkan.",
			Data:          nil,
		}
	}

	return result
}

// UpdateJurusan melakukan pembaruan data di tabel Keuangan berdasarkan ID.
func UpdateHusadainance(DP DataPerusahaan.Keuangan) models.BaseResponseModels {
	var query string
	var result models.BaseResponseModels
	db := connection.DB

	// Query SQL untuk memperbarui data di tabel Keuangan berdasarkan ID
	query = "UPDATE Keuangan SET Hari = ?, Transaksi_Masuk = ?, Transaksi_Keluar = ?, Total = ?, Bulan = ?, Jumlah_Transaksi = ?, Total_Transaksi = ?, Keterangan = ? WHERE Id = ?"

	// Melakukan eksekusi query dengan data yang diberikan
	tempResult := db.Exec(query, DP.Hari, DP.Transaksi_Masuk, DP.Transaksi_Keluar, DP.Total, DP.Bulan, DP.Jumlah_Transaksi, DP.Total_Transaksi, DP.Keterangan, DP.Id)

	// Memeriksa apakah terdapat kesalahan selama eksekusi query
	if tempResult.Error != nil {
		result = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Memeriksa apakah ada baris yang terpengaruh oleh perintah UPDATE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil diubah.",
				Data:          nil,
			}
		}
	}

	return result
}

// DeleteJurusan menghapus data di tabel Keuangan berdasarkan ID.
func DeleteHusadaVinance(request DataPerusahaan.Keuangan) models.BaseResponseModels {
	var query string
	var result models.BaseResponseModels
	db := connection.DB
	query = "DELETE FROM Keuangan WHERE Id = ?"

	tempResult := db.Exec(query, request.Id)

	if tempResult.Error != nil {
		result = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Periksa apakah ada baris yang terpengaruh oleh perintah DELETE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil dihapus.",
				Data:          nil,
			}
		}
	}

	return result
}

func GetHusadaVinanceByID(request DataPerusahaan.Keuangan) models.BaseResponseModels {
	var query string
	var result models.BaseResponseModels
	var DataPerusahaan []DataPerusahaan.Keuangan
	db := connection.DB

	if request.Id != 0 {
		query = "SELECT * FROM Keuangan WHERE Id = ?"
	} else {
		query = "SELECT * FROM Keuangan"
	}

	tempResult := db.Raw(query).Find(& DataPerusahaan)

	if tempResult.Error != nil {
		result = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {	
		result = models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:         	DataPerusahaan,
		}
	}

	return result
}
