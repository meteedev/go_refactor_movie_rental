package rental

const REGULAR = 0
const NEW_RELEASE = 1
const CHILDRENS = 2

type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) (rcvr Movie) {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}

func (m Movie) PriceCode() int {
	return m.priceCode
}
func (m Movie) Title() string {
	return m.title
}
