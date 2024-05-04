package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) (c Customer) {
	return Customer{
		rentals: make([]Rental, 0),
		name:    name,
	}
}

func (c Customer) AddRental(arg Rental) {
	c.rentals = append(c.rentals, arg)
}
func (c Customer) Name() string {
	return c.name
}

func amountFor(each Rental) (amount float64){
	result := 0.0
	switch each.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if each.DaysRented() > 2 {
			result += float64(each.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(each.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if each.DaysRented() > 3 {
			result += float64(each.DaysRented()-3) * 1.5
		}
	}
	return result
}


func (c Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	//result := fmt.Sprintf("%v%v%v", "Rental Record for ", c.GetName(), "\n")
	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, each := range c.rentals {
		
		thisAmount := amountFor(each)

		frequentRenterPoints++
		if each.Movie().PriceCode() == NEW_RELEASE && each.DaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("\t%v\t%.1f\n", each.Movie().Title(), thisAmount)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints )
	return result
}
