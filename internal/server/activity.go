package server

import (
	"fmt"
	"sync"
	"time"
)

var ErrIDNotFound = fmt.Errorf("ID not found")

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          uint64    `json:"id"`
}

type Activities struct {
	mu         sync.Mutex
	activities []Activity
}

func (c *Activities) Insert(activity Activity) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	activity.ID = uint64(len(c.activities) + 1)
	c.activities = append(c.activities, activity)
	return activity.ID
}

func (c *Activities) Retrieve(id uint64) (Activity, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if id > uint64(len(c.activities)) {
		return Activity{}, ErrIDNotFound
	}
	return c.activities[id-1], nil
}
