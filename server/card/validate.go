package card

func ValidateCardNumber(cardNumber int64) (bool, error) {
	return (cardNumber%10+checksum(cardNumber/10))%10 == 0, nil
}

func checksum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		current := number % 10

		if i%2 == 0 {
			current = current * 2
			if current > 9 {
				current = current%10 + current/10
			}
		}

		luhn += current
		number = number / 10
	}

	return luhn % 10
}
