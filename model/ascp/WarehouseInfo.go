package ascp

// WarehouseInfo 结构体
type WarehouseInfo struct {
	// 联系人信息；必填 1个
	ContactInfos []ContactInfo `json:"contact_infos,omitempty" xml:"contact_infos>contact_info,omitempty"`
	// 仓库编码，string（50)    卖家下唯一主键 erp维度是会员ID+仓库ID生成同步过来的仓库编码，确保唯一，接口交互时货主下资源的唯一主键；
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
	// 商家编码，string（50)   商家在erp设置的商家编码，商家可编辑
	ErpWarehouseBizCode string `json:"erp_warehouse_biz_code,omitempty" xml:"erp_warehouse_biz_code,omitempty"`
	// 仓库名称，string（50）
	ErpWarehouseName string `json:"erp_warehouse_name,omitempty" xml:"erp_warehouse_name,omitempty"`
	// WMS仓编码，string（50）
	WmsWarehouseCode string `json:"wms_warehouse_code,omitempty" xml:"wms_warehouse_code,omitempty"`
	// WMS仓名称，string（50）
	WmsWarehouseName string `json:"wms_warehouse_name,omitempty" xml:"wms_warehouse_name,omitempty"`
	// 省份，string（15)
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市，string（15）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 地区，string（15）
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 乡镇，string（15）
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址，string（50）
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 邮编，string（15）
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 状态：0=停用；1=启用
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 服务商自定义的仓编码，服务商+业务身份下唯一
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 服务商自定义的仓库名称，创建时必填
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// WMS系统仓code，创建时必填
	WmsStoreCode string `json:"wms_store_code,omitempty" xml:"wms_store_code,omitempty"`
	// WMSAppkey，创建时必填
	WmsAppkey string `json:"wms_appkey,omitempty" xml:"wms_appkey,omitempty"`
	// 仓拓展信息，创建时必填
	ExtendInfo *ExtendInfo `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
}
