package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func (t testing.TB, wallet Wallet, want Cash) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func (t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("Deposit and view Balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Cash(10))
		assertBalance(t, wallet, Cash(10))
	})

	t.Run("Withdraw and view Balance", func(t *testing.T) {
		wallet := Wallet{balance: Cash(10)}
		wallet.Withdraw(Cash(10))
		assertBalance(t, wallet, Cash(0))
	})

	t.Run("Overwithdraw", func(t *testing.T) {
		wallet := Wallet{balance: Cash(10)}
		err := wallet.Withdraw(Cash(100))
		assertBalance(t, wallet, Cash(10))

		assertError(t, err)
	})
}