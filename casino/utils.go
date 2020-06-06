package main

func (c *casinoServer) hasEnoughTokens(usrID userID, target int32) bool {
	return c.userToTokens[usrID] >= target
}

func (c *casinoServer) hasEnoughStocks(usrID userID, stocks int32) bool {
	return c.userToStocks[usrID] >= stocks
}
