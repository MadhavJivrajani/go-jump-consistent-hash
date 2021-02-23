package main

var multiplier uint64 = 2862933555777941757
var one int64 = 1

// Hash implements the jump consistent algorithm in O(1) memory
// and linear time. It takes in the key to be hashed of type
// unit64 and the number of buckets available of type int32, and
// returns a bucket in the range [0, numBuckets)
// Ref: https://arxiv.org/pdf/1406.2294.pdf
func Hash(key uint64, numBuckets int32) int64 {
	var bucket int64 = -1
	var j int64 = 0
	for j < int64(numBuckets) {
		bucket = j
		key = key*multiplier + 1
		var factor float64 = float64(one<<31) / float64((key>>33)+1)
		j = int64(float64((bucket + 1)) * factor)
	}
	return bucket
}
