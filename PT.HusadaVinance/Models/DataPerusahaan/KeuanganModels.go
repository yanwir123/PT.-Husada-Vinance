package dataperusahaan

type Keuangan struct {
	Id               int     `json:"id" from:"id"`
	Hari             string  `json:"hari" from:"hari"`
	Transaksi_Masuk  float64 `json:"Transaksi Masuk" from:"Transaksi Masuk"`
	Transaksi_Keluar float64 `json:"Transaksi Keluar" from:"Transaksi Keluar"`
	Total            float64 `json:"Total" from:"Total"`
	Bulan            string  `json:"Bulan" from:"Bulan"`
	Jumlah_Transaksi float64 `json:"Jumlah Transaksi" from:"Jumlah Transaksi"`
	Total_Transaksi  float64 `json:"Total Transaksi" from:"Total Transaksi"`
	Keterangan       string  `json:"Keterangan" from:"Keterangan"`
}