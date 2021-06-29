package tmalltrend

// TrendStyleBindInfoBO 
type TrendStyleBindInfoBO struct {
    // 款式关联趋势词信息列表，趋势词信息：趋势词id+叶子类目id，单个款式最多关联3个趋势词
    TrendWordInfoList   []Null `json:"trend_word_info_list,omitempty" xml:"trend_word_info_list>null,omitempty"`
    // 同步操作目的，枚举，INSERT("新增"),     UPDATE("更新"),     OFFLINE("下线");
    SyncPurpose   string `json:"sync_purpose,omitempty" xml:"sync_purpose,omitempty"`
    // 款式编号，业务唯一
    StyleSerialNumber   string `json:"style_serial_number,omitempty" xml:"style_serial_number,omitempty"`
}
