package user

import (
	"github.com/itsyouonline/identityserver/db"
)

type DigitalAssetAddress struct {
	CurrencySymbol string      `json:"currencysymbol"`
	Address        string      `json:"address"`
	Label          string      `json:"label"`
	Expire         db.DateTime `json:"expire"`
	Noexpiration	 bool				 `json:"noexpiration"`
}
