package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	pr := float64(productionRate) / 100 // I got an error here because I was doing float64 type 
    // conversion for the whole operation i.e. float64(productionRate / 100) that led to an error
    // for multiple test cases. This is because integer division messed things up.
    // float64(productionRate / 100) has both operands being int so the Go compiler does integer 
    // division first.
    // So if productionRate was 221 then 221 by 100 would be 2, which would then be converted to 
    // float64 type becoming 2.00 . However, if I did float64(productionRate) / 100, I would get 2.21
    // which would be correct. Because untyped constants adopt types based on the context. So 
    // multiplying by 0.01 or dividing by 100 works because the operand they are interacting, in 
    // this case, productionRate is of type floating point.
    sr := float64(successRate) * 0.01
    wr := pr * sr
    return wr * 100 // Apparently you can multiply a float variable like wr with 100 but if I were to 
    // declare a variable called percent of type int with a value 100 and and multiply it I get an
    // error. This is due to the same reason as above wherein untyped constants get typed based on c
    // context.
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
/*
Mistakes I made earlier and why they were wrong:

1) Integer division at the wrong time
   - I had: carsPerMin := productionRate / 60
   - Both operands were int, so Go did integer division first (truncation). Example: 221/60 == 3.
     That lost the fractional cars/minute and skewed every subsequent calculation.

2) Treating a percentage as a whole number instead of a fraction
   - I multiplied by successRate directly (e.g., 90) instead of using 0.90.
   - Converting with int(successRate) or float64(successRate) alone is not enough; I must divide by 100
     (or multiply by 0.01) to convert a percent to its fractional form.

3) Type-mixing confusion (constants vs. typed values)
   - In Go, untyped numeric constants (like 60 or 100) adopt the type of the other operand.
     If both sides are int, integer math happens; if one side is float64, the operation is floating-point.
     This is why float64(productionRate)/60 works (60 becomes 60.0), but productionRate/60 converted
     afterwards would already be truncated.

Why this version works:

1) float64 before dividing by 60.0
   - carsPerMin := float64(productionRate) / 60.0
   - Forcing float math preserves the fractional part (no truncation).

2) Convert percent to fraction
   - successRate is given in [0,100]; multiplying by 0.01 converts it to [0,1].
   - minSucc := carsPerMin * (successRate * 0.01)

3) Truncate only once at the end
   - The spec/tests want an int count per minute. Returning int(minSucc) truncates toward zero at the final step.

Sanity checks:
- (221, 100%)  → (221/60) ≈ 3.683 → 3
- (426, 80%)   → (426*0.8)/60 = 5.68 → 5
- (6824, 20.5%)→ (6824*0.205)/60 ≈ 23.31 → 23
*/
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	
    carsPerMin := float64(productionRate)/60.0
    minSucc := (float64(carsPerMin) * (float64(successRate) * 0.01))
    return int(minSucc)
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    carsGroup := carsCount/10
    carsSingle := carsCount%10
    return uint(95000*carsGroup) + uint(10000*carsSingle)
}