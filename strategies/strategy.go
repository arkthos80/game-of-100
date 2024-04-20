package strategies

type Strategy int

const (
	Random        Strategy = iota //choose a random movement
	CrossFirst                    //first try go horizontal or vertical (randomly), use Diagonal only if needed (randomly)
	DiagonalFirst                 //first try go Diagonal (randomly), use Cross only if needed (randomly)
	//RandomHVFirstDiagInOrder          //unimplemented
	//RandomDiagFirstHVInOrder          //unimplemented
	//StayCloseToBorder                 //unimplemented
)

var All = []Strategy{Random, CrossFirst, DiagonalFirst}

var AllNames = [...]string{
	"Random",
	"CrossFirst",
	"DiagonalFirst",
}

// GetStrategyByName returns the enum Strategy corresponding to the given name.
// If the name is not found, it returns -1.
func GetStrategyByName(name string) Strategy {
	for i, n := range AllNames {
		if n == name {
			return All[i]
		}
	}
	return -1
}
