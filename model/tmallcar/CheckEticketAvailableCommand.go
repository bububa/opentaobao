package tmallcar

import (
	"sync"
)

// CheckEticketAvailableCommand 结构体
type CheckEticketAvailableCommand struct {
	// 核销码
	EticketCode string `json:"eticket_code,omitempty" xml:"eticket_code,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolCheckEticketAvailableCommand = sync.Pool{
	New: func() any {
		return new(CheckEticketAvailableCommand)
	},
}

// GetCheckEticketAvailableCommand() 从对象池中获取CheckEticketAvailableCommand
func GetCheckEticketAvailableCommand() *CheckEticketAvailableCommand {
	return poolCheckEticketAvailableCommand.Get().(*CheckEticketAvailableCommand)
}

// ReleaseCheckEticketAvailableCommand 释放CheckEticketAvailableCommand
func ReleaseCheckEticketAvailableCommand(v *CheckEticketAvailableCommand) {
	v.EticketCode = ""
	v.StoreId = 0
	poolCheckEticketAvailableCommand.Put(v)
}
