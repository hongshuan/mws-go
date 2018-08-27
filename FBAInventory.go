package mws

/**
 * Get Service Status
 * Gets the status of the service.
 *     Status is one of GREEN, RED representing:
 *     GREEN: The service section is operating normally.
 *     RED: The service section disruption.
 *
 * @param mixed request array of parameters for GetServiceStatus request or GetServiceStatus object itself
 * @see FBAInventoryServiceMWS_Model_GetServiceStatusRequest
 * @return GetServiceStatusResponse
 *
 * @throws FBAInventoryException
 */
func GetServiceStatus(request) {
}

/**
 * List Inventory Supply
 * Get information about the supply of seller-owned inventory in
 *     Amazon's fulfillment network. "Supply" is inventory that is available
 *     for fulfilling (a.k.a. Multi-Channel Fulfillment) orders. In general
 *     this includes all sellable inventory that has been received by Amazon,
 *     that is not reserved for existing orders or for internal FC processes,
 *     and also inventory expected to be received from inbound shipments.
 *
 *     This operation provides 2 typical usages by setting different
 *     ListInventorySupplyRequest value:
 *
 *     1. Set value to SellerSkus and not set value to QueryStartDateTime,
 *     this operation will return all sellable inventory that has been received
 *     by Amazon's fulfillment network for these SellerSkus.
 *
 *     2. Not set value to SellerSkus and set value to QueryStartDateTime,
 *     This operation will return information about the supply of all seller-owned
 *     inventory in Amazon's fulfillment network, for inventory items that may have had
 *     recent changes in inventory levels. It provides the most efficient mechanism
 *     for clients to maintain local copies of inventory supply data.
 *
 *     Only 1 of these 2 parameters (SellerSkus and QueryStartDateTime) can be set value for 1 request.
 *     If both with values or neither with values, an exception will be thrown.
 *
 *     This operation is used with ListInventorySupplyByNextToken
 *     to paginate over the resultset. Begin pagination by invoking the
 *     ListInventorySupply operation, and retrieve the first set of
 *     results. If more results are available,continuing iteratively requesting further
 *     pages results by invoking the ListInventorySupplyByNextToken operation (each time
 *     passing in the NextToken value from the previous result), until the returned NextToken
 *     is null, indicating no further results are available.
 *
 * @param mixed request array of parameters for ListInventorySupply request or ListInventorySupply object itself
 * @see FBAInventoryServiceMWS_Model_ListInventorySupplyRequest
 * @return FBAInventoryServiceMWS_Model_ListInventorySupplyResponse
 *
 * @throws FBAInventoryServiceMWS_Exception
 */
func ListInventorySupply(request) {
}

/**
 * List Inventory Supply By Next Token
 * Continues pagination over a resultset of inventory data for inventory
 *     items.
 *
 *     This operation is used in conjunction with ListUpdatedInventorySupply.
 *     Please refer to documentation for that operation for further details.
 *
 * @param mixed request array of parameters for ListInventorySupplyByNextToken request or ListInventorySupplyByNextToken object itself
 * @see ListInventorySupplyByNextTokenRequest
 * @return ListInventorySupplyByNextTokenResponse
 *
 * @throws FBAInventoryServiceMWS_Exception
 */
func ListInventorySupplyByNextToken(request) {
}
