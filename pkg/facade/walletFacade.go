package facade

import (
	"errors"
	"fmt"
	"sort"
)

type walletFacade struct {
	accounts     Accounts
	wallets      Wallets
	securityCode *securityCode
	notification *notification
}

//CreateNewWallet is create new wallet
func (w *walletFacade) CreateNewWallet(balance *int, owner string) (*Wallet, error) {

	account, exist := w.accounts.searchAccount(nil, &owner)
	if !exist {
		newAccount, err := w.accounts.createNewAccount(owner)
		if err != nil {
			return nil, err
		}

		if balance != nil {
			newAccount.amount = *balance
		}

		account = newAccount

		w.accounts = append(w.accounts, *account)

	}

	newWallet, err := w.wallets.createNewWallet(balance, &account.ID)
	if err != nil {
		return nil, err
	}

	w.wallets = append(w.wallets, *newWallet)

	_ = newWallet.printNewWalletInfo()

	sort.Sort(Accounts(w.accounts))
	sort.Sort(Wallets(w.wallets))
	return newWallet, nil
}

func (w *walletFacade) PutMoney(walletID int, secureCode int, amount int) error {
	wallet, exist := w.wallets.searchWallet(walletID)

	if !exist {
		return fmt.Errorf("Wallet with id %v doesn't exist", walletID)
	}

	acc, exist := w.accounts.searchAccount(&wallet.accountID, nil)
	if !exist {
		return errors.New("You doesn't have an account")
	}

	err := w.accounts.putMoney(&acc.ID, amount)
	if err != nil {
		return err
	}

	err = w.wallets.putMoney(wallet.id, amount)
	if err != nil {
		return err
	}

	fmt.Println("Money added to wallet !!!!")
	return nil
}

func (w *walletFacade) GetMoney(walletID int, amount int) error {
	wallet, exist := w.wallets.searchWallet(walletID)
	if !exist {
		return fmt.Errorf("Wallet with id %v doesn't exist", walletID)
	}

	acc, exist := w.accounts.searchAccount(&wallet.accountID, nil)
	if !exist {
		return errors.New("You doesn't have an account")
	}

	if amount > acc.amount {
		return errors.New("account :: too much")
	}

	acc.amount -= amount

	if amount > wallet.balance {
		return errors.New("wallet :: too much")
	}

	wallet.balance -= amount

	return nil
}

func (w *walletFacade) ShowFacadeSustem() {

	fmt.Println(w.accounts)
	fmt.Println(w.wallets)

}

func CreateWalletSystem() *walletFacade {

	fmt.Printf("Creating new wallets table ....")

	walletSystem := &walletFacade{
		accounts:     createNewAccountsStorage(),
		wallets:      newWalletTable(),
		securityCode: &securityCode{},
		notification: &notification{},
	}

	fmt.Println("done")

	return walletSystem
}
