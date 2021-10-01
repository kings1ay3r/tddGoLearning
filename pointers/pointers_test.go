package pointers

import "testing"

func TestPointer(t *testing.T) {

	assertTest := func(t testing.TB, got, want string) {
		if got != want {
			errormsg := "Expecting " + want + " and recieved " + got
			t.Errorf(errormsg)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		if got == nil {
			t.Fatal("Wanted an error didnt get one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Withdrawls", func(t *testing.T) {
		wallet := Wallet{33}
		err := wallet.Withdraw(10)

		got := wallet.Balance()
		want := "USD 23,000"
		assertTest(t, got, want)
		if err != nil {
			t.Error("Error")
		}
	})
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := "USD 10,000"
		assertTest(t, got, want)
	})
	t.Run("Insufficient Balance", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(100)

		got := wallet.Balance()
		want := "USD 20,000"
		assertTest(t, got, want)
		assertError(t, err, InsufficientFundsError)
	})
}
