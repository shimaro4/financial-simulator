package family

type Income struct {
	Earner string
	Source string
	Amount float64
	YearlyAvgIncrease float64
}

type Expense struct {
	Housing float64
	HousingType string
	Groceries float64
	Utilities float64
	Miscl float64
}

type Savings struct {
	KidsEducation float64
	Entertainment float64
	RainyDay float64
	RepairMaint float64
}

type Investment struct {
	MonthlyContributions float64
	TotalInvested float64
	Balance float64
	ReturnRate float64
}
