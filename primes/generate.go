package primes

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

// Send the sequence 2, 3, 4, … to channel 'ch'.
func seededGenerate(stop chan bool, ch chan<- int) {
	go func() {
		defer close(ch)
		for i := 2; ; i++ {
			select {
			case <-stop:
				return
			case ch <- i:
			}
		}
	}()
}

//GenerateNNumberOfPrimes generates N number of primes and gives them back in a list.
func GenerateNNumberOfPrimes(numberOfPrimes int) []int {
	var primes []int
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a subprocess.
	for i := 0; i < numberOfPrimes; i++ {
		prime := <-ch
		primes = append(primes, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	return primes
}

//GeneratePrimesToALimit Generates primes up until a certain limit.
func GeneratePrimesToALimit(limit int) []int {
	var primes []int
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a subprocess.
	for {
		prime := <-ch
		if prime > limit {
			return primes
		}
		primes = append(primes, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

//SeedPrimeNumbers Generates primes up until a certain limit.
func SeedPrimeNumbers(stop chan bool, limit int) <-chan int {
	out := make(chan int)
	ch := make(chan int) // Create a new channel.
	go seededGenerate(stop, ch)      // Start generate() as a subprocess.
	go func() {
		defer close(out)
		for {
			prime := <-ch
			if prime > limit {
				out <- prime
				return
			}
			out <- prime
			ch1 := make(chan int)
			go filter(ch, ch1, prime)
			ch = ch1
		}
	}()
	return out
}
