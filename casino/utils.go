package main

func (c *casinoServer) hasEnoughTokens(usrID userID, target int32) bool {
	return c.userToTokens[usrID] >= target
}

func (c *casinoServer) canBuy(usrID userID, price int32) bool {
	return c.userToTokens[usrID] >= price
}

func (c *casinoServer) canSell(usrID userID, stocks int32) bool {
	return c.userToStocks[usrID] > stocks
}
