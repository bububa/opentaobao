package icbulogistics

// WarehouseDto 
type WarehouseDto struct {
    // 备注
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 工作时间
    WorkingTime   string `json:"working_time,omitempty" xml:"working_time,omitempty"`
    // 联系电话
    ContactPhone   string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
    // 联系人
    ContactPerson   string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
    // 地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 仓库名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 仓库编码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
