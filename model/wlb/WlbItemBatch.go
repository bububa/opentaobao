package wlb

// WlbItemBatch 
type WlbItemBatch struct {
    // 商品批次记录id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 用户id
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 商品数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 残次数量
    DefectQuantity   int64 `json:"defect_quantity,omitempty" xml:"defect_quantity,omitempty"`
    // 存储类型
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 批次编号
    BatchCode   string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
    // 生产编号
    ProduceCode   string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
    // 到期时间
    DueDate   string `json:"due_date,omitempty" xml:"due_date,omitempty"`
    // 生产日期
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    // 接受日期
    ReceiveDate   string `json:"receive_date,omitempty" xml:"receive_date,omitempty"`
    // 保质期
    GuaranteePeriod   string `json:"guarantee_period,omitempty" xml:"guarantee_period,omitempty"`
    // 天（单位）
    GuaranteeUnit   int64 `json:"guarantee_unit,omitempty" xml:"guarantee_unit,omitempty"`
    // 产地
    ProduceArea   string `json:"produce_area,omitempty" xml:"produce_area,omitempty"`
    // 描述
    Remarks   string `json:"remarks,omitempty" xml:"remarks,omitempty"`
    // 状态。item_batch_status_open:开放 item_batch_status_lock:冻结 item_batch_status_invalid:无效
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 创建者
    Creator   string `json:"creator,omitempty" xml:"creator,omitempty"`
    // 最后修改者
    LastModifier   string `json:"last_modifier,omitempty" xml:"last_modifier,omitempty"`
    // 版本
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 最后修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 是否删除。0：正常 1：删除
    IsDeleted   bool `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}
