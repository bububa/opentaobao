package wdk

// ReceiptBatchInfo 
type ReceiptBatchInfo struct {
    // 批次名称
    BatchName   string `json:"batch_name,omitempty" xml:"batch_name,omitempty"`
    // 批次号
    BatchId   string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
    // 履约单集合
    FulfillOrderList   []FulfillOrder `json:"fulfill_order_list,omitempty" xml:"fulfill_order_list>fulfill_order,omitempty"`
    // 容器信息
    ContainerInfoList   []string `json:"container_info_list,omitempty" xml:"container_info_list>string,omitempty"`
    // 波次标识
    BcFlag   string `json:"bc_flag,omitempty" xml:"bc_flag,omitempty"`
    // 容器数量
    ContainerCount   int64 `json:"container_count,omitempty" xml:"container_count,omitempty"`
    // 扩展属性
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
}
