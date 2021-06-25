package wdk

// ConveyorBasicConfigDTO 
type ConveyorBasicConfigDTO struct {

    // 滑道数目
    SlidewayCount   int64 `json:"slideway_count,omitempty"`

    // 异常滑道ID
    ExceptionSlidewayId   int64 `json:"exception_slideway_id,omitempty"`

    // 批次任务超时时间（分钟）
    BatchTaskTimeoutInterval   int64 `json:"batch_task_timeout_interval,omitempty"`

    // 是否启用阻挡器
    EnableSeparator   bool `json:"enable_separator,omitempty"`

    // 是否启用监控
    EnableMonitor   bool `json:"enable_monitor,omitempty"`

    // 单件通道亮三色灯
    EnableSingleSlidewayLight   bool `json:"enable_single_slideway_light,omitempty"`

    // 最大转圈数目
    MaxOrbitingNum   int64 `json:"max_orbiting_num,omitempty"`

    // 分配策略
    DispatchSlidewayPolicy   int64 `json:"dispatch_slideway_policy,omitempty"`

    // 容器识别器1编号
    ContainerReaderBox1   string `json:"container_reader_box1,omitempty"`

    // 容器识别器2编号
    ContainerReaderBox2   string `json:"container_reader_box2,omitempty"`

    // 悬挂链控制盒子编号
    ConveyorBox   string `json:"conveyor_box,omitempty"`

    // 滑道分组个数
    SlidewayGroupCount   int64 `json:"slideway_group_count,omitempty"`

}
