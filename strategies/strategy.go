package strategies

type Strategy int

const (
	Random                   Strategy = iota
	CrossFirstRandom                  //first try go horizontal or vertical (randomly), use Diagonal only if needed (randomly)
	RandomHVFirstDiagInOrder          //unimplemented
	DiagonalFirstRandom               //first try go Diagonal (randomly), use Cross only if needed (randomly)
	RandomDiagFirstHVInOrder          //unimplemented
	StayCloseToBorder                 //unimplemented
)

var All = []Strategy{Random, CrossFirstRandom, RandomHVFirstDiagInOrder, DiagonalFirstRandom, RandomDiagFirstHVInOrder, StayCloseToBorder}

var AllNames = [...]string{
	"Random",
	"CrossFirstRandom",
	"RandomHVFirstDiagInOrder",
	"DiagonalFirstRandom",
	"RandomDiagFirstHVInOrder",
	"StayCloseToBorder",
}
