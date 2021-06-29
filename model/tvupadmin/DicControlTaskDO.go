package tvupadmin

// DicControlTaskDO 
type DicControlTaskDO struct {
    // 任务类型
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 操作者
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    // 任务描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 任务名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // uuid
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    // apk id
    ApkId   int64 `json:"apk_id,omitempty" xml:"apk_id,omitempty"`
    // 设备型号
    Devices   string `json:"devices,omitempty" xml:"devices,omitempty"`
    // 牌照方
    License   int64 `json:"license,omitempty" xml:"license,omitempty"`
}
