package task

import "time"

type StartFn func()
type StopFn func()
type UpdateFn func()

type TaskerOption func(*TaskerOptions)
type TaskerOptions struct {
	uniqueId       int32
	startFns       []StartFn // start callback
	stopFns        []StopFn  // task stop callback
	updateFn       UpdateFn  // default update callback
	timer          *time.Timer
	d              time.Duration // timeout duration
	updateInterval time.Duration // update interval duration
	executeTimeout time.Duration // execute timeout
	chanBufSize    int           // channel buffer size
}

func defaultTaskerOptions() *TaskerOptions {
	return &TaskerOptions{
		uniqueId:       0,
		d:              TaskDefaultTimeout,
		startFns:       make([]StartFn, 0, 5),
		stopFns:        make([]StopFn, 0, 5),
		updateFn:       nil,
		timer:          time.NewTimer(TaskDefaultTimeout),
		updateInterval: TaskDefaultUpdateInterval,
		executeTimeout: TaskDefaultExecuteTimeout,
		chanBufSize:    TaskDefaultChannelSize,
	}
}

func WithStartFns(f ...StartFn) TaskerOption {
	return func(o *TaskerOptions) {
		o.startFns = o.startFns[:0]
		o.startFns = append(o.startFns, f...)
	}
}

func WithStopFns(f ...StopFn) TaskerOption {
	return func(o *TaskerOptions) {
		o.stopFns = o.stopFns[:0]
		o.stopFns = append(o.stopFns, f...)
	}
}

func WithUpdateFn(f UpdateFn) TaskerOption {
	return func(o *TaskerOptions) {
		o.updateFn = f
	}
}

func WithTimeout(d time.Duration) TaskerOption {
	return func(o *TaskerOptions) {
		o.d = d
	}
}

func WithUpdateInterval(d time.Duration) TaskerOption {
	return func(o *TaskerOptions) {
		o.updateInterval = d
	}
}

func WithExecuteTimeout(d time.Duration) TaskerOption {
	return func(o *TaskerOptions) {
		o.executeTimeout = d
	}
}

func WithChannelBufferSize(sz int) TaskerOption {
	return func(o *TaskerOptions) {
		o.chanBufSize = sz
	}
}

func WithUniqueId(id int32) TaskerOption {
	return func(o *TaskerOptions) {
		o.uniqueId = id
	}
}
