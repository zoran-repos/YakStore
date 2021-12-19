package functions

func GetTotalMilk(elapsedTime int, oldInAges float64) float64 {
	var milkCount float64
	yakLife := int(oldInAges * 100)
	for i := 0; i < elapsedTime; i++ {
		milkCount += GetMilk(yakLife + i)
	}
	return milkCount
}

func GetMilk(yaklife int) float64 {
	return 50 - (float64(yaklife) * 0.03)
}

func GetSkin(elapsedTime int, oldInAges float64) int {
	skinCount := 1
	ageLastShave := int(oldInAges * 100)
	allowedGapInShave := 8 + float64(ageLastShave)*0.01
	for day := 1; day < elapsedTime; day++ {
		currentAgeInDays := int(oldInAges*100 + float64(day))
		if int(oldInAges)*100+day < 1000 {
			if (currentAgeInDays - ageLastShave) > int(allowedGapInShave) {
				skinCount++
				ageLastShave = currentAgeInDays
			}

		}
	}
	return skinCount
}
