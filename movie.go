package rental

const REGULAR 		= 0
const NEW_RELEASE 	= 1
const CHILDRENS 	= 2



type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) (rcvr Movie) {
	return Movie{
		title: title,
		priceCode: priceCode,
	}
}

func (rcvr Movie) PriceCode() int {
	return rcvr.priceCode
}
func (rcvr Movie) GetTitle() string {
	return rcvr.title
}
func (rcvr Movie) SetPriceCode(arg int) {
	rcvr.priceCode = arg
}
