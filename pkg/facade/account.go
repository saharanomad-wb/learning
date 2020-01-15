package facade

import (
	"errors"
	"fmt"
	"math/rand"
)

//account base type for account
type account struct {
	ID     int
	Owner  string
	wallet int
	amount int
}

//Accounts is accounts storage
type Accounts []account

func (a Accounts) Len() int           { return len(a) }
func (a Accounts) Less(i, j int) bool { return a[i].ID < a[j].ID }
func (a Accounts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a Accounts) createNewAccount(owner string) (*account, error) {
	if owner == "" {
		return nil, errors.New("Owner name can't be empty")
	}

	_, exist := a.searchAccount(nil, &owner)

	if !exist {
		newAccount := account{
			ID:    generateID(),
			Owner: owner,
		}
		a = append(a, newAccount)

		return &newAccount, nil
	}

	return nil, fmt.Errorf("Account with name %v exist", owner)

}

func (a Accounts) searchAccount(accountID *int, owner *string) (*account, bool) {
	if len(a) == 0 {
		return nil, false
	}

	if accountID == nil && owner == nil {
		return nil, false
	}

	if accountID != nil && owner == nil {
		left := 0
		right := len(a) - 1

		for left <= right {
			mid := left + ((right - left) / 2)
			if a[mid].ID == *accountID {
				return &a[mid], true
			}
			if a[mid].ID > *accountID {
				right = mid - 1
			}

			if a[mid].ID < *accountID {
				left = mid + 1
			}
		}
		return nil, false
	}

	if owner != nil && accountID == nil {
		for i := 0; i < len(a); i++ {
			if a[i].Owner == *owner {
				return &a[i], true
			}
		}
		return nil, false
	}

	return nil, false

}

func (a Accounts) putMoney(accountID *int, amount int) error {

	if accountID == nil {
		return errors.New("You have to provide account id")
	}

	if amount == 0 {
		return errors.New("amount have to be greater than 0")
	}

	acc, exist := a.searchAccount(accountID, nil)
	if !exist {
		return fmt.Errorf("Account with id %v doesn't exist")
	}

	acc.amount += amount

	return nil
}

func (a Accounts) getMoney(accountID *int, amount int) (*int, error) {
	if accountID == nil {
		return nil, errors.New("You have to provide account id")
	}

	if amount == 0 {
		return nil, errors.New("amount have to be greater than 0")
	}

	acc, exist := a.searchAccount(accountID, nil)
	if !exist {
		return nil, fmt.Errorf("Account with id %v doesn't exist")
	}

	if amount > acc.amount {
		return nil, fmt.Errorf("Request amount is exceed your balance")
	}

	acc.amount -= amount

	return &amount, nil
}

func generateID() int {
	return rand.Intn(1000)
}

func createNewAccountsStorage() Accounts {
	fmt.Printf("Creating accounts storage .... ")
	storage := Accounts{}
	fmt.Println("done")
	return storage
}
