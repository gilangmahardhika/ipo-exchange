package model

import (
	"time"

	"github.com/lib/pq"
)

type Ipo struct {
	ID                    string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Code                  string         `gorm:"type:varchar(10);Index" json:"code"`
	Name                  string         `gprm:"type:varchar(300)" json:"name"`
	EipoID                int            `gorm:"Index" json:"eipo_id"`
	SharedLot             int            `json:"shared_lot"`
	Logo                  string         `gorm:"type:varchar(255)" json:"logo"`
	FinalPrice            int            `json:"final_price"`
	Status                string         `gorm:"index;type:varchar(255)" json:"status"`
	Sector                string         `gorm:"index;type:varchar(255)" json:"sector"`
	SubSector             string         `gorm:"type:varchar(255)" json:"sub_sector"`
	LineOfBusiness        string         `gorm:"type:varchar(300)" json:"line_of_business"`
	CompanyOverview       string         `json:"company_overview"`
	Address               string         `json:"address"`
	Website               string         `gorm:"type:varchar(255)" json:"website"`
	PercentageOfShare     int            `json:"percentage_of_share"`
	ParticipantAdmin      string         `gorm:"type:varchar(255)" json:"participant_admin"`
	OfferingStartDate     time.Time      `gorm:"index" json:"offering_start_date"`
	OfferingEndDate       time.Time      `gorm:"index" json:"offering_end_date"`
	OfferingPrice         int            `json:"offering_price"`
	AllotmentDate         time.Time      `gorm:"index" json:"allotment_date"`
	DistributionDate      time.Time      `gorm:"index" json:"distribution_date"`
	ListingDate           time.Time      `gorm:"index" json:"listing_date"`
	AdditionalInformation string         `gorm:"type:varchar(255)" json:"additional_information"`
	BookBuildingStartDate time.Time      `gorm:"index" json:"book_building_start_date"`
	BookBuildingEndDate   time.Time      `gorm:"index" json:"book_building_end_date"`
	BookBuildingMinPrice  int            `json:"book_building_min_price"`
	BookBuildingMaxPrice  int            `json:"book_building_max_price"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	Prospectus            pq.StringArray `gorm:"type:varchar(300)[]" json:"prospectus"`
	Underwriters          pq.StringArray `gorm:"type:varchar(300)[]" json:"underwriters"`
}
