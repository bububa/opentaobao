package qimen

// WarehouseInfos 
type WarehouseInfos struct {
    // 仓库编码，string（50）
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    // 仓库名称，string（50）
    WarehouseName   string `json:"warehouseName,omitempty" xml:"warehouseName,omitempty"`
    // 省份，string（15）
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // 城市，string（15）
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 地区，string（15）
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 乡镇，string（15）
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    // 详细地址，string（50）
    DetailAddress   string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
    // 仓库名称，string（50）
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 仓库电话，string（20）
    Tel   string `json:"tel,omitempty" xml:"tel,omitempty"`
    // 负责人手机，string（20）
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 仓库状态，string（20）
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
}
