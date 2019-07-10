package main

import (
	"testing"
)

//v4
func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

//v3
// func TestWallet(t *testing.T) {

// 	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
// 		t.Helper()
// 		got := wallet.Balance()

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}

// 	assertError := func(t *testing.T, got error, want string) {
// 		t.Helper()
// 		if got == nil {
// 			t.Fatal("didn't get an error but wanted one")
// 		}

// 		if got.Error() != want {
// 			t.Errorf("got '%s', want '%s'", got, want)
// 		}
// 	}

// 	t.Run("Deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))
// 		assertBalance(t, wallet, Bitcoin(10))

// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}
// 		wallet.Withdraw(Bitcoin(10))
// 		assertBalance(t, wallet, Bitcoin(10))
// 	})

// 	t.Run("Withdraw insufficient funds", func(t *testing.T) {
// 		startingBalance := Bitcoin(20)
// 		wallet := Wallet{startingBalance}
// 		err := wallet.Withdraw(Bitcoin(100))

// 		assertBalance(t, wallet, startingBalance)
// 		assertError(t, err, "cannot withdraw, insufficient funds")
// 	})

// }

//v2
// func TestWallet(t *testing.T) {

// 	t.Run("Deposit", func(t *testing.T) {
// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}

// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}

// 		wallet.Withdraw(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// }

//v1
// func TestWallet(t *testing.T) {

// 	wallet := Wallet{}

// 	wallet.Deposit(Bitcoin(10))

// 	got := wallet.Balance()

// 	want := Bitcoin(10)

// 	if got != want {
// 		t.Errorf("got %s want %s", got, want)
// 	}

// }
