package algo

type (
	// Edge is an interface that must be implemented by the types passed to Kruskal
	Edge interface {
		GetSource() string
		GetDestination() string
		GetWeight() float64
	}
)
