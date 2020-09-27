package pyfuncs

// returns channel and is a generator
func rangeGenerator(startVal float64, endVal float64, delta float64) <-chan float64 {
	var (
		counter = startVal
		retChan = make(chan float64)
	)
	go func() {
		// Check if adding the delta wont lead to infinite loop
		// increase range
		defer close(retChan)
		if startVal <= endVal && delta > 0 {
			for counter <= endVal {
				retChan <- counter
				counter = counter + delta
			}
			// decrease range
		} else if startVal >= endVal && delta < 0 {
			for counter >= endVal {
				retChan <- counter
				counter = counter + delta
			}
		}
	}()
	return retChan
}

// Range from python
func Range(valuesIn ...float64) <-chan float64 {
	var (
		startVal   float64
		endVal     float64
		delta      float64
		nrOfParams int = len(valuesIn)
		retChan        = make(chan float64)
	)

	// If none parameters where given
	if nrOfParams == 0 {
		close(retChan)
		return retChan

		// Range from 0 to the endVal, step 1 or -1
	} else if nrOfParams == 1 {
		endVal = valuesIn[0]
		startVal = 0
		// set delta so it will achive the goal
		if endVal > 0 {
			delta = 1
		} else {
			delta = -1
		}
		return rangeGenerator(startVal, endVal, delta)

		// Range from the startVal to endVal, step 1 or -1
	} else if nrOfParams == 2 {
		endVal = valuesIn[1]
		startVal = valuesIn[0]
		// set delta so it will achive goal
		if endVal-startVal > 0 {
			delta = 1
		} else {
			delta = -1
		}
		return rangeGenerator(startVal, endVal, delta)

		// Range from startVal to endVal, step is delta
	} else if nrOfParams == 3 {
		endVal = valuesIn[1]
		startVal = valuesIn[0]
		delta = valuesIn[2]
		return rangeGenerator(startVal, endVal, delta)
	} else {
		close(retChan)
		return retChan
	}
}
