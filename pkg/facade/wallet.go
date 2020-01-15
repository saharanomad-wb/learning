package facade

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

const (
	randRange = 100000000
)

type wallet struct {
	id         int
	balance    int
	secureCode int
	accountID  int
}

type wallets []wallet

func (w wallets) Len() int           { return len(w) }
func (w wallets) Less(i, j int) bool { return w[i].id < w[j].id }
func (w wallets) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }

func newWalletTable() wallets {
	return wallets{}
}

func (w wallets) createNewWallet(balance *int, accountID *int) (*wallet, error) {

	if accountID == nil {
		return nil, errors.New("You have to provide account id")
	}

	initialBalance := 0

	if balance != nil {
		initialBalance = *balance
	}

	t := Wallet{
		id:         generateID(),
		balance:    initialBalance,
		secureCode: rand.Intn(randRange),
		accountID:  *accountID,
	}

	w = append(w, t)

	sort.Sort(wallets(w))

	return &t, nil
}

func (w wallets) putMoney(id int, amount int) error {
	wallet, exist := w.searchWallet(id)
	if !exist {
		return fmt.Errorf("Wallet with id %v doesn't exist", id)
	}

	wallet.balance += amount

	return nil

}

func (w wallets) getMoney(id int, amount int) error {
	wallet, exist := w.searchWallet(id)
	if !exist {
		return fmt.Errorf("Wallet with id %v doesn't exist", id)
	}
	if amount > wallet.balance {
		return fmt.Errorf("too much")
	}

	wallet.balance -= amount
	return nil
}

func (w wallets) searchWallet(id int) (*Wallet, bool) {

	left := 0
	right := len(w) - 1

	for left <= right {
		mid := left + ((right - left) / 2)
		if w[mid].id == id {
			return &w[mid], true
		}
		if w[mid].id > id {
			right = mid - 1
		}

		if w[mid].id < id {
			left = mid + 1
		}
	}
	return nil, false
}

func (wl *Wallet) printNewWalletInfo() error {
	fmt.Printf("Created :: wallet id: %v wallet balance: %v wallet secureCode: %v wallet account: %v \n", wl.id, wl.balance, wl.secureCode, wl.accountID)
	return nil
}
