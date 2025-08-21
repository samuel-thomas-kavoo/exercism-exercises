package interest
import "fmt"

const(
    negBalance = 3.213
    smolBalance = 0.5
    medBalance = 1.621
    bigBalance = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
    case balance < 0:
        return float32(negBalance)
    case balance >=0 && balance < 1000:
        return float32(smolBalance)
    case balance >=1000 && balance < 5000:
        return float32(medBalance)
	case balance >=5000:
        return float32(bigBalance)
    default:
        fmt.Println("default case triggered. Oops.")
        return 0.0 //Even though logically the switch covers all possible values of balance, the Go compiler requires an explicit return. It doesn’t “trust” that one of the case blocks will always match. This is a safety feature.
// The default case initially only called fmt.Println(...) but did not return a value. This was not allowed by Go as the compiler requires that all control paths (branches, switches, etc.) either return a value of the correct type or end with a panic().
    }
    return 0.0 //Adding this here too as a default fallback return for example's sake. 
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    switch {
    case balance < 0:
        interest := (balance*negBalance)/float64(100)
        return interest
    case balance >=0 && balance < 1000:
        interest := (balance*smolBalance)/float64(100)
        return interest
    case balance >=1000 && balance < 5000:
        interest := (balance*medBalance)/float64(100)
        return interest
	case balance >=5000:
        interest := (balance*bigBalance)/float64(100)
        return interest
    default:
        fmt.Println("default case triggered. Oops.")
        return 0.0
    }
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate (balance float64) float64 {
    switch {
    case balance < 0:
    	interest := (balance*negBalance)/float64(100)
        balance += interest
        return balance
    case balance >=0 && balance < 1000:
        interest := (balance*smolBalance)/float64(100)
        balance += interest
        return balance
    case balance >=1000 && balance < 5000:
        interest := (balance*medBalance)/float64(100)
        balance += interest
        return balance
	case balance >=5000:
        interest := (balance*bigBalance)/float64(100)
        balance += interest
        return balance
    }
    return 0.0
}

const epsilon = 0.00001
// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance (balance, targetBalance float64) int {
    var yearsLeft int = 0
    if balance == targetBalance {
        return yearsLeft
    } else if balance > targetBalance {
        return yearsLeft
    } else {
        for i := balance; i < targetBalance - epsilon; {
            switch {
                case balance < 0:
                interest := (balance*negBalance)/float64(100)
                balance += interest
                yearsLeft++
                i = balance
                case balance >=0 && balance < 1000:
                interest := (balance*smolBalance)/float64(100)
                balance += interest
                yearsLeft++
                i = balance
                case balance >=1000 && balance < 5000:
                interest := (balance*medBalance)/float64(100)
                balance += interest
                yearsLeft++
                i = balance
                case balance >=5000:
                interest := (balance*bigBalance)/float64(100)
                balance += interest
                yearsLeft++
                i = balance
            }
    	}
    }
return yearsLeft
}
