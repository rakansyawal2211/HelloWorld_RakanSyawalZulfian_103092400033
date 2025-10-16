package main

import (
	"fmt"
	"strings"
)

type User struct {
	Email    string
	Password string
	Balance  float64
}

var users [10]User
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

	if strings.TrimSpace(email) == "" {
		fmt.Println("Email tidak boleh kosong.")
		return
	}

	if strings.Contains(email, " ") {
		fmt.Println("Email tidak boleh mengandung spasi.")
		return
	}

	for i := 0; i < 10; i++ {
		if users[i].Email == email {
			fmt.Println("Email sudah terdaftar.")
			return
		}
	}

	fmt.Print("Masukkan password (tanpa spasi): ")
	fmt.Scanln(&password)

	if strings.TrimSpace(password) == "" {
		fmt.Println("Password tidak boleh kosong.")
		return
	}

	if strings.Contains(password, " ") {
		fmt.Println("Password tidak boleh mengandung spasi.")
		return
	}

	for i := 0; i < 10; i++ {
		if users[i].Email == "" {
			users[i] = User{Email: email, Password: password, Balance: 1000}
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

	if email == "" || password == "" {
		fmt.Println("Error: Email atau password tidak boleh kosong.")
		return
	}

	for i := 0; i < 10; i++ {
		if users[i].Email == email && users[i].Password == password {
			currentUser = &users[i]
			fmt.Printf("Login berhasil! Selamat datang, %s\n", currentUser.Email)
			mainMenu()
			return
		}
	}

	fmt.Println("Email atau password salah.")
}

func mainMenu() {
	for {
		var choice int
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Transfer Uang")
		fmt.Println("3. Top Up")
		fmt.Println("4. Logout")
		fmt.Print("Masukkan pilihan (1/2/3/4): ")
		fmt.Scanln(&choice)

		if choice == 1 {
			checkBalance()
		} else if choice == 2 {
			transferMoney()
		} else if choice == 3 {
			topUp()
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
	var emailTujuan string
	var jumlah float64

	fmt.Print("\nMasukkan email tujuan transfer: ")
	fmt.Scanln(&emailTujuan)

	if emailTujuan == currentUser.Email {
		fmt.Println("Tidak bisa transfer ke diri sendiri.")
		return
	}

	var targetUser *User = nil
	for i := 0; i < 10; i++ {
		if users[i].Email == emailTujuan {
			targetUser = &users[i]
			break
		}
	}

	if targetUser == nil {
		fmt.Println("Pengguna dengan email tersebut tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan jumlah yang ingin ditransfer: Rp")
	fmt.Scanln(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Jumlah tidak valid.")
		return
	}

	if currentUser.Balance < jumlah {
		fmt.Println("Saldo Anda tidak cukup.")
		return
	}

	currentUser.Balance -= jumlah
	targetUser.Balance += jumlah

	fmt.Printf("Transfer berhasil ke %s sebesar Rp%.2f\n", targetUser.Email, jumlah)
}

func topUp() {
	var amount float64
	var method, confirm string

	fmt.Println("\n== Top Up E-Money ==")
	fmt.Print("Masukkan jumlah top up: Rp")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Jumlah tidak valid.")
		return
	}

	fmt.Println("Pilih metode pembayaran:")
	fmt.Println("1. Bank Transfer")
	fmt.Println("2. E-Wallet")
	fmt.Println("3. Virtual Account")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scanln(&method)

	if method != "1" && method != "2" && method != "3" {
		fmt.Println("Metode tidak valid.")
		return
	}

	fmt.Print("Ketik 'konfirmasi' untuk melanjutkan pembayaran: ")
	fmt.Scanln(&confirm)

	if confirm != "konfirmasi" {
		fmt.Println("Top up dibatalkan.")
		return
	}

	var metodePembayaran string
	if method == "1" {
		metodePembayaran = "Bank Transfer"
	} else if method == "2" {
		metodePembayaran = "E-Wallet"
	} else {
		metodePembayaran = "Virtual Account"
	}

	currentUser.Balance += amount
	fmt.Printf("Top up sebesar Rp%.2f melalui %s berhasil.\n", amount, metodePembayaran)
	fmt.Printf("Saldo Anda sekarang: Rp%.2f\n", currentUser.Balance)
}

func logout() {
	currentUser = nil
	fmt.Println("Anda telah logout.")
}
