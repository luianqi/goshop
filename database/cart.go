package database

import "errors"

func AddProductToCart() {

}

// custom errors
var (
	ErrCannotFindProduct    = errors.New("Cannot find the product")
	ErrCannotDecodeProducts = errors.New("Cannot find the product")
	ErrUserIdIsNotValid     = errors.New("User is not valid")
	ErrCannotUpdateUser     = errors.New("Cannot add this product to the cart")
	ErrCannotRemoveItem     = errors.New("Cannot remove this item from the cart")
	ErrCannotGetItem        = errors.New("Was unable to get item from the cart")
	ErrCannotBuyCartItem    = errors.New("Cannot update the purchase")
)

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
