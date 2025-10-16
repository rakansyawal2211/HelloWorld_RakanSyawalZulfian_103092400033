package main

import "fmt"

type User struct {
	Email    string
	Password string
	Balance  float64 // Menambahkan saldo pengguna
}

var users [10]User // Array dengan panjang tetap untuk menyimpan data pengguna
var currentUser *User

func main() {
	fmt.Println("=== Selamat datang di Aplikasi E-Money ===")

	for {
		var choice int
		fmt.Println("\nPilih opsi:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan pilihan (1/2/3): ")
		fmt.Scanln(&choice)

		if choice == 1 {
			login()
		} else if choice == 2 {
			register()
		} else if choice == 3 {
			fmt.Println("Terima kasih, sampai jumpa!")
			break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func register() {
	var email, password string
	fmt.Println("\n== Registrasi ==")
	fmt.Print("Masukkan email (tanpa spasi): ")
	fmt.Scanln(&email)

	// Cek apakah email sudah terdaftar dalam array
	for i := 0; i < 10; i++ { // Kita tahu ada 10 slot di array
		if users[i].Email == email {
			fmt.Println("Email sudah terdaftar.")
			return
		}
	}

	// Meminta password pengguna
	fmt.Print("Masukkan password (tanpa spasi): ")
	fmt.Scanln(&password)

	// Cari posisi kosong di array untuk menyimpan user baru
	for i := 0; i < 10; i++ { // Kita tahu ada 10 slot di array
		if users[i].Email == "" { // Posisi kosong ditemukan
			users[i] = User{Email: email, Password: password, Balance: 1000} // Menambahkan saldo awal
			fmt.Println("Registrasi berhasil! Silakan login.")
			return
		}
	}
	fmt.Println("Registrasi gagal! Kapasitas pengguna sudah penuh.")
}

func login() {
	var email, password string
	fmt.Println("\n== Login ==")
	fmt.Print("Masukkan email: ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	// Cek jika email atau password kosong
	if email == "" || password == "" {
		fmt.Println("Error: Email atau password tidak boleh kosong.")
		return
	}

	// Cek apakah ada pengguna terdaftar
	userFound := false

	// Cari user berdasarkan email dan password
	for i := 0; i < 10; i++ { // Kita tahu ada 10 slot di array
		if users[i].Email == email && users[i].Password == password {
			currentUser = &users[i] // Gunakan indeks array untuk menyimpan pointer ke user
			fmt.Printf("Login berhasil! Selamat datang, %s\n", currentUser.Email)
			userFound = true
			// Lanjutkan ke menu utama
			mainMenu()
			break
		}
	}

	// Jika tidak ada pengguna ditemukan, login gagal
	if !userFound {
		fmt.Println("Email atau password salah.")
	}
}

func mainMenu() {
	for {
		var choice int
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Transfer Uang")
		fmt.Println("3. Lihat Riwayat Transaksi")
		fmt.Println("4. Logout")
		fmt.Print("Masukkan pilihan (1/2/3/4): ")
		fmt.Scanln(&choice)

		if choice == 1 {
			checkBalance()
		} else if choice == 2 {
			transferMoney()
		} else if choice == 3 {
			viewTransactionHistory()
		} else if choice == 4 {
			logout()
			break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func checkBalance() {
	fmt.Printf("\nSaldo Anda: Rp%.2f\n", currentUser.Balance)
}

func transferMoney() {
	var amount float64
	fmt.Print("\nMasukkan jumlah yang ingin ditransfer: Rp")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Jumlah transfer tidak valid.")
		return
	}

	if currentUser.Balance < amount {
		fmt.Println("Saldo tidak cukup.")
		return
	}

	// Transfer berhasil
	currentUser.Balance -= amount
	fmt.Printf("Transfer berhasil! Saldo Anda kini Rp%.2f\n", currentUser.Balance)
}

func viewTransactionHistory() {
	// Dalam contoh ini, tidak ada riwayat transaksi yang disimpan
	fmt.Println("\nFitur riwayat transaksi belum tersedia.")
}

func logout() {
	currentUser = nil
	fmt.Println("Anda telah berhasil logout.")
}
