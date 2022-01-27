package import

import "rsc.io/quote"

func ImportQuote() string {
return quote.Go() + "\n" + quote.Glass();
}
