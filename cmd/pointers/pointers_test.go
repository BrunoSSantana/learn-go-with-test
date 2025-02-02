package pointers_test

import (
	"testing"

	pointers "github.com/BrunoSSantana/learn-go-with-test/cmd/pointers"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))

		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := pointers.Wallet{}

		wallet.Deposit(pointers.Bitcoin(20))
		err := wallet.Withdraw(pointers.Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := pointers.Bitcoin(20)

		wallet := pointers.Wallet{}
		wallet.Deposit(startingBalance)
		err := wallet.Withdraw(pointers.Bitcoin(100))

		assertError(t, err, pointers.ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet pointers.Wallet, want pointers.Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got  an error but didn't want one")
	}
}
