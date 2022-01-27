package drit

import "rsc.io/quote"

func DritQuote() string {
return quote.Go() + "\n" + quote.Glass();
}
func Drit2Quote() string {
return quote.Opt() +"\n" + quote.Hello();
}
