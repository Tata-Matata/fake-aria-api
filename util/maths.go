package util

func BytesToTiB(bytes uint64) float64 {
	//1 TiB in bytes; Binary (base 2, 1 TiB = 2⁴⁰ bytes)
	//Equivalent to multiplying by 2ⁿ where n is 40
	const TiB = 1 << 40
	return float64(bytes) / float64(TiB)
}
