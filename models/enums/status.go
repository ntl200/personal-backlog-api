package enums

type Status int64
const (
	NotYetStarted	Status = iota + 1
	OnGoing
	Completed 		
)

func (s Status) String() string {
	return [...]string{"Not yet started", "On going", "Completed"}[s-1]
}

func (s Status) EnumIndex() int {
	return int(s)
}