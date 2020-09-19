package stats

import (
	"github.com/Iftikhor99/bankAverageMain/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		types.Payment {
			Amount: 10_000_00,
		},
		types.Payment {
			Amount: 20_000_00,
		},
		types.Payment {
			Amount: 30_000_00,
		},
	}
	fmt.Println(Avg(payments))
	//Output:2000000
	
}
