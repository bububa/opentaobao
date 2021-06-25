package wdk

// SyncVersionBo 
type SyncVersionBo struct {

    // 外部唯一ID
    OutUniqueId   string `json:"out_unique_id,omitempty"`

    // 数据周期结束时间
    RangeEndTime   int64 `json:"range_end_time,omitempty"`

    // 数据周期开始时间
    RangeStartTime   int64 `json:"range_start_time,omitempty"`

    // 同步数据总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 同步的数据表名称
    TableName   string `json:"table_name,omitempty"`

    // 操作ID
    OperateId   string `json:"operate_id,omitempty"`

    // WDK_MARKET--营销数据
    BizCode   string `json:"biz_code,omitempty"`

    // 版本号
    VersionId   int64 `json:"version_id,omitempty"`

}
