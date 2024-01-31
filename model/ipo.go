package model

import "time"

type Ipo struct {
	Code                  string    `json:"code"`
	Name                  string    `json:"name"`
	EipoID                uint32    `json:"eipo_id"`
	SharedLot             uint64    `json:"shared_lot"`
	Logo                  string    `json:"logo"`
	FinalPrice            uint32    `json:"final_price"`
	Status                string    `json:"status"`
	Sector                string    `json:"sector"`
	SubSector             string    `json:"sub_sector"`
	LineOfBusiness        string    `json:"line_of_business"`
	CompanyOverview       string    `json:"company_overview"`
	Address               string    `json:"address"`
	Website               string    `json:"website"`
	PercentageOfShare     uint32    `json:"percentage_of_share"`
	ParticipantAdmin      string    `json:"participant_admin"`
	Underwriters          []string  `json:"underwriters"`
	OfferingStartDate     time.Time `json:"offering_start_date"`
	OfferingEndDate       time.Time `json:"offering_end_date"`
	OfferingPrice         uint32    `json:"offering_price"`
	AllotmentDate         time.Time `json:"allotment_date"`
	DistributionDate      time.Time `json:"distribution_date"`
	ListingDate           time.Time `json:"listing_date"`
	Prospectus            []string  `json:"pospectus"`
	AdditionalInformation string    `json:"additional_information"`
	BookBuildingStartDate time.Time `json:"book_building_start_date"`
	BookBuildingEndDate   time.Time `json:"book_building_end_date"`
	BookBuildingMinPrice  uint64    `json:"book_building_min_price"`
	BookBuildingMaxPrice  uint64    `json:"book_building_max_price"`
}
