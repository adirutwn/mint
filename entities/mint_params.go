package entities

type MintParams struct {
	AssetCode string
	Amount string
	IssuerAddress string
	IssuerSeedKey string
	IsIssuerAccountLock bool
	DistributorAddress string
	DistributorSeedKey string
	IsDistributorAccountNeedToCreate bool
	IsDistributorAccountNeedToTrust bool
}
