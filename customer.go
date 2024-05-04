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

type Record struct{
	Name string
	TotalAmount float64
	Points int
	Charges  []Charge
}

type Charge struct{
	Title string
	Amount float64
}

func renderPlainText( r Record)string{
	result := fmt.Sprintf("Rental Record for %v\n", r.Name)

	for _, c := range r.Charges {
		result += fmt.Sprintf("\t%v\t%.1f\n", c.Title, c.Amount)
	}

	result += fmt.Sprintf("Amount owed is %.1f\n", r.TotalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", r.Points)

	return result
}



func (c Customer) Statement() string {
	
	totalAmount := getTotalAmount(c.rentals)
	frequentRenterPoints := getTotalPoints(c.rentals)
	
	charges := []Charge{}
	for _, r := range c.rentals {
		charge := Charge{
			Title: r.Movie().Title(),
			Amount: r.Charge(),
		}
		charges = append(charges, charge)
	}

	record := Record{
		Name: c.Name(),
		TotalAmount: totalAmount,
		Points: frequentRenterPoints,
		Charges: charges,
	}
	return renderPlainText(record)
}