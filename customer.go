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
	frequentRenterPoints :=0
	frequentRenterPoints++
		if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
			frequentRenterPoints++
		}
	return frequentRenterPoints
}


func (c Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	//result := fmt.Sprintf("%v%v%v", "Rental Record for ", c.GetName(), "\n")
	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, r := range c.rentals {
		
		frequentRenterPoints += getPoints(r)
		
		result += fmt.Sprintf("\t%v\t%.1f\n", r.Movie().Title(), r.Charge())
		totalAmount += r.Charge()
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints )
	return result
}
