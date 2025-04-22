package lentil

// IBeanstalk is an interface to alllow queue DI and mocking.
type IBeanstalk interface {
	Watch(tube string) (int, error)
	Ignore(tube string) (int, error)
	Use(tube string) error
	Put(priority, delay, ttr int, data []byte) (uint64, error)
	ReserveWithTimeout(seconds int) (*Job, error)
	Delete(id uint64) error
	Release(id uint64, pri, delay int) error
	Reserve() (*Job, error)
	Bury(id uint64, pri int) error
	Touch(id uint64) error
	Peek(id uint64) (*Job, error)
	PeekReady() (*Job, error)
	PeekDelayed() (*Job, error)
	PeekBuried() (*Job, error)
	StatsJob(id uint64) (map[string]string, error)
	StatsTube(tube string) (map[string]string, error)
	Stats() (map[string]string, error)
	Kick(bound int) (int, error)
	ListTubes() ([]string, error)
	ListTubeUsed() (string, error)
	ListTubesWatched() ([]string, error)
	Quit() error
	PauseTube(tube string, delay int) error
}

type ILentil interface {
	Dial(addr string, readerSize int, args ...interface{}) (IBeanstalk, error)
}

type Lentil struct{}

func (Lentil) Dial(addr string, readerSize int, args ...interface{}) (IBeanstalk, error) {
	q, err := Dial(addr, readerSize, args...)
	if err != nil {
		return nil, err
	}
	return q, nil
}
