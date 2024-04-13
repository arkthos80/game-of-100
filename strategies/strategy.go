package strategies

type Strategy int

const (
	Random                   Strategy = iota
	RandomHVFirstDiagRandom           //first try go horizontal or vertical, use Diagonal only if needed. Diagonal move will be random
	RandomHVFirstDiagInOrder          //unimplemented
	RandomDiagFirstHVRandom           //unimplemented
	RandomDiagFirstHVInOrder          //unimplemented
	StayCloseToBorder                 //unimplemented
)

var strategies = []Strategy{Random, RandomHVFirstDiagRandom, RandomHVFirstDiagInOrder, RandomDiagFirstHVRandom, RandomDiagFirstHVInOrder, StayCloseToBorder}
