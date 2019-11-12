package structs

// ShipResponse returned data from /ship
type ShipResponse struct {
	StatusCode    int
	StatusMessage string
	Message       string
	Ship          Ship
}

// ShipsResponse returns data from /ships
type ShipsResponse struct {
	StatusCode    int
	StatusMessage string
	Message       string
	Ships         []SmallShip
}

// ConstructionResponse returns data from /build
type ConstructionResponse struct {
	StatusCode    int
	StatusMessage string
	Message       string
	Construction  Construction
}
