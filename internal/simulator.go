package internal

func (h *Household) AddIncome(earner, source string, amount, yearlyIncrease float64){
	newIncome := Income {
		Earner: earner,
		Source: source,
		Amount: amount,
		YearlyAvgIncrease: yearlyIncrease,
	}

	h.Incomes = append(h.Incomes, newIncome)
}