package eventbus

import "time"

type Event struct {
	Type      string                 
	Timestamp time.Time              
	Source    string                 
	Payload   map[string]interface{}
}
