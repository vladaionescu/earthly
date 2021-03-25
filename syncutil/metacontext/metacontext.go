package metacontext

import (
	"context"
	"sync"
	"time"
)

var _ context.Context = &MetaContext{}

// MetaContext is an object which implements context.Context and which holds multiple
// contexts within it. The MetaContext is considered canceled only when ALL of the
// underlying contexts have been canceled.
type MetaContext struct {
	subDoneCh chan int // index

	mu           sync.Mutex
	doneCh       chan struct{}
	numDone      int
	firstDoneErr error
	sub          []context.Context
}

// New returns a new metacontext.
func New(ctx context.Context) *MetaContext {
	mc := &MetaContext{
		doneCh:    make(chan struct{}),
		subDoneCh: make(chan int),
	}
	mc.Add(ctx)
	go mc.monitor()
	return mc
}

func (mc *MetaContext) monitor() {
	for index := range mc.subDoneCh {
		mc.mu.Lock()
		mc.numDone++
		if mc.numDone == 1 {
			mc.firstDoneErr = mc.sub[index].Err()
		}
		if mc.numDone == len(mc.sub) {
			close(mc.doneCh)
		}
		mc.mu.Unlock()
	}
}

// Add adds a new context to the metacontext.
func (mc *MetaContext) Add(ctx context.Context) error {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	if mc.isDone() {
		return mc.Err()
	}
	mc.sub = append(mc.sub, ctx)
	index := len(mc.sub) - 1
	go func() {
		<-ctx.Done()
		mc.subDoneCh <- index
	}()
	return nil
}

// Deadline returns the Deadline of the fist context.
func (mc *MetaContext) Deadline() (deadline time.Time, ok bool) {
	mc.mu.Lock()
	if len(mc.sub) == 0 {
		mc.mu.Unlock()
		return time.Time{}, false
	}
	ctx := mc.sub[0]
	mc.mu.Unlock()
	return ctx.Deadline()
}

// Done returns the done channel. The MetaContext is done only when ALL of the
// contained contexts are done.
func (mc *MetaContext) Done() <-chan struct{} {
	return mc.doneCh
}

// Err returns the first done error reported by any context.
func (mc *MetaContext) Err() error {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	return mc.firstDoneErr
}

// Value calls context.Value on the first not-done context, or on the first one,
// if all are done.
func (mc *MetaContext) Value(key interface{}) interface{} {
	mc.mu.Lock()
	if len(mc.sub) == 0 {
		mc.mu.Unlock()
		return nil
	}
	// Find the first not-done ctx. If none found, use the first one.
	var selectedCtx context.Context
	for _, ctx := range mc.sub {
		select {
		case <-mc.doneCh:
			continue
		default:
		}
		selectedCtx = ctx
		break
	}
	if selectedCtx == nil {
		selectedCtx = mc.sub[0]
	}
	mc.mu.Unlock()
	return selectedCtx.Value(key)
}

func (mc *MetaContext) isDone() bool {
	select {
	case <-mc.doneCh:
		return true
	default:
		return false
	}
}
