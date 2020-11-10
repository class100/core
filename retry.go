package core

import (
	`time`
)

type Retry struct {
	// Count 重试次数
	Count int
	// WaitTime 重试等待时间（每两次重试之间的间隔时间）
	WaitTime time.Duration
	// MaxWaitTime 最大重试等待时间
	MaxWaitTime time.Duration
}
