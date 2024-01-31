package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ipo struct {
	gorm.Model
	ID                    string          `gorm:"primaryKey" json:"id"`
	Code                  string          `gorm:"uniqueIndex" json:"code"`
	Name                  string          `json:"name"`
	EipoID                uint32          `json:"eipo_id"`
	SharedLot             uint64          `json:"shared_lot"`
	Logo                  string          `json:"logo"`
	FinalPrice            *uint32         `json:"final_price"`
	Status                string          `gorm:"index" json:"status"`
	Sector                *string         `gorm:"index" json:"sector"`
	SubSector             *string         `json:"sub_sector"`
	LineOfBusiness        *string         `json:"line_of_business"`
	CompanyOverview       *string         `json:"company_overview"`
	Address               *string         `json:"address"`
	Website               *string         `json:"website"`
	PercentageOfShare     *uint32         `json:"percentage_of_share"`
	ParticipantAdmin      *string         `json:"participant_admin"`
	Underwriters          *pq.StringArray `gorm:"type:varchar(64)[]" json:"underwriters"`
	OfferingStartDate     *time.Time      `gorm:"index" json:"offering_start_date"`
	OfferingEndDate       *time.Time      `gorm:"index" json:"offering_end_date"`
	OfferingPrice         *uint32         `json:"offering_price"`
	AllotmentDate         *time.Time      `gorm:"index" json:"allotment_date"`
	DistributionDate      *time.Time      `gorm:"index" json:"distribution_date"`
	ListingDate           *time.Time      `gorm:"index" json:"listing_date"`
	Prospectus            *pq.StringArray `gorm:"type:varchar(64)[]" json:"pospectus"`
	AdditionalInformation *string         `json:"additional_information"`
	BookBuildingStartDate *time.Time      `gorm:"index" json:"book_building_start_date"`
	BookBuildingEndDate   *time.Time      `gorm:"index" json:"book_building_end_date"`
	BookBuildingMinPrice  *uint64         `json:"book_building_min_price"`
	BookBuildingMaxPrice  *uint64         `json:"book_building_max_price"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
}
