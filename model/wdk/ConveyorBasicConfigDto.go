package wdk

import (
	"sync"
)

// ConveyorBasicConfigDto 结构体
type ConveyorBasicConfigDto struct {
	// 容器识别器1编号
	ContainerReaderBox1 string `json:"container_reader_box1,omitempty" xml:"container_reader_box1,omitempty"`
	// 容器识别器2编号
	ContainerReaderBox2 string `json:"container_reader_box2,omitempty" xml:"container_reader_box2,omitempty"`
	// 悬挂链控制盒子编号
	ConveyorBox string `json:"conveyor_box,omitempty" xml:"conveyor_box,omitempty"`
	// 滑道数目
	SlidewayCount int64 `json:"slideway_count,omitempty" xml:"slideway_count,omitempty"`
	// 异常滑道ID
	ExceptionSlidewayId int64 `json:"exception_slideway_id,omitempty" xml:"exception_slideway_id,omitempty"`
	// 批次任务超时时间（分钟）
	BatchTaskTimeoutInterval int64 `json:"batch_task_timeout_interval,omitempty" xml:"batch_task_timeout_interval,omitempty"`
	// 最大转圈数目
	MaxOrbitingNum int64 `json:"max_orbiting_num,omitempty" xml:"max_orbiting_num,omitempty"`
	// 分配策略
	DispatchSlidewayPolicy int64 `json:"dispatch_slideway_policy,omitempty" xml:"dispatch_slideway_policy,omitempty"`
	// 滑道分组个数
	SlidewayGroupCount int64 `json:"slideway_group_count,omitempty" xml:"slideway_group_count,omitempty"`
	// 是否启用阻挡器
	EnableSeparator bool `json:"enable_separator,omitempty" xml:"enable_separator,omitempty"`
	// 是否启用监控
	EnableMonitor bool `json:"enable_monitor,omitempty" xml:"enable_monitor,omitempty"`
	// 单件通道亮三色灯
	EnableSingleSlidewayLight bool `json:"enable_single_slideway_light,omitempty" xml:"enable_single_slideway_light,omitempty"`
}

var poolConveyorBasicConfigDto = sync.Pool{
	New: func() any {
		return new(ConveyorBasicConfigDto)
	},
}

// GetConveyorBasicConfigDto() 从对象池中获取ConveyorBasicConfigDto
func GetConveyorBasicConfigDto() *ConveyorBasicConfigDto {
	return poolConveyorBasicConfigDto.Get().(*ConveyorBasicConfigDto)
}

// ReleaseConveyorBasicConfigDto 释放ConveyorBasicConfigDto
func ReleaseConveyorBasicConfigDto(v *ConveyorBasicConfigDto) {
	v.ContainerReaderBox1 = ""
	v.ContainerReaderBox2 = ""
	v.ConveyorBox = ""
	v.SlidewayCount = 0
	v.ExceptionSlidewayId = 0
	v.BatchTaskTimeoutInterval = 0
	v.MaxOrbitingNum = 0
	v.DispatchSlidewayPolicy = 0
	v.SlidewayGroupCount = 0
	v.EnableSeparator = false
	v.EnableMonitor = false
	v.EnableSingleSlidewayLight = false
	poolConveyorBasicConfigDto.Put(v)
}
