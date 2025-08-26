package family

type Household struct {
	Name string
	NumPeople int
	State string
	Incomes []Income
	Expense Expense
	Savings Savings
	Investment Investment
}

func (h *Household) AddIncome(earner, source string, amount, yearlyIncrease float64){
	newIncome := Income {
		Earner: earner,
		Source: source,
		Amount: amount,
		YearlyAvgIncrease: yearlyIncrease,
	}

	h.Incomes = append(h.Incomes, newIncome)
}
