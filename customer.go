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

func (r Rental) Charge() (amount float64){
	result := 0.0
	switch r.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if r.DaysRented() > 2 {
			result += float64(r.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if r.DaysRented() > 3 {
			result += float64(r.DaysRented()-3) * 1.5
		}
	}
	return result
}

func getPoints(r Rental)(points int){	
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func getTotalAmount(rentals []Rental)(totalAmount float64){
	result := 0.0
	for _, r := range rentals {
		result += r.Charge()
	}
	return result
}

func getTotalPoints(rentals []Rental)(getTotalPoints int){
	result := 0
	for _, r := range rentals {
		result += getPoints(r)
	}
	return result
}

func (c Customer) Statement() string {
	
	totalAmount := getTotalAmount(c.rentals)
	frequentRenterPoints := getTotalPoints(c.rentals)
	
	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, r := range c.rentals {
		result += fmt.Sprintf("\t%v\t%.1f\n", r.Movie().Title(), r.Charge())
	}

	



	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints )
	return result
}
