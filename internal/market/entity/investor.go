package entity

type Investor struct {
	Id            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		Id:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

type InvestorAssetPosition struct {
	AssetId string
	Shares  int
}
