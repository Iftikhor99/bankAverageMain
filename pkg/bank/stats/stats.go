package stats

import 	"github.com/Iftikhor99/bankAverageMain/pkg/bank/types"

//Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	average := types.Money(0)
	length := len(payments)
	for _, payment := range payments {
		average += payment.Amount	
	}
	
	average = average / types.Money(length)
	return average
} 