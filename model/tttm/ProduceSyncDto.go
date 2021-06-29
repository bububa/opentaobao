package tttm

// ProduceSyncDTO 
type ProduceSyncDTO struct {
    // 顺序
    LinkSort   int64 `json:"link_sort,omitempty" xml:"link_sort,omitempty"`
    // 生产节点
    ProduceLink   string `json:"produce_link,omitempty" xml:"produce_link,omitempty"`
    // 产量
    ProduceNum   int64 `json:"produce_num,omitempty" xml:"produce_num,omitempty"`
    // 次品量
    DefectiveNum   int64 `json:"defective_num,omitempty" xml:"defective_num,omitempty"`
    // 生产状态
    ProduceStatus   int64 `json:"produce_status,omitempty" xml:"produce_status,omitempty"`
    // 提交时间
    SubmitTime   string `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
    // 实际时间
    FinishTime   string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
}
