package rental

const(
	_ =iota
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) (m Movie) {
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
