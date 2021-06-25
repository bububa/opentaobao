package qimen

// WarehouseInfos 
type WarehouseInfos struct {

    // 仓库编码，string（50）
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 仓库名称，string（50）
    WarehouseName   string `json:"warehouseName,omitempty"`

    // 省份，string（15）
    Province   string `json:"province,omitempty"`

    // 城市，string（15）
    City   string `json:"city,omitempty"`

    // 地区，string（15）
    Area   string `json:"area,omitempty"`

    // 乡镇，string（15）
    Town   string `json:"town,omitempty"`

    // 详细地址，string（50）
    DetailAddress   string `json:"detailAddress,omitempty"`

    // 仓库名称，string（50）
    Name   string `json:"name,omitempty"`

    // 仓库电话，string（20）
    Tel   string `json:"tel,omitempty"`

    // 负责人手机，string（20）
    Mobile   string `json:"mobile,omitempty"`

    // 仓库状态，string（20）
    Status   string `json:"status,omitempty"`

}