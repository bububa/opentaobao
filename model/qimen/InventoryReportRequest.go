package qimen

import (
	"sync"
)

// InventoryReportRequest 结构体
type InventoryReportRequest struct {
	// 商品库存信息列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 盘点单编码
	CheckOrderCode string `json:"checkOrderCode,omitempty" xml:"checkOrderCode,omitempty"`
	// 仓储系统的盘点单编码
	CheckOrderId string `json:"checkOrderId,omitempty" xml:"checkOrderId,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 盘点时间(YYYY-MM-DD HH:MM:SS)
	CheckTime string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
	OutBizCode string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 变动类型：CHECK=盘点 ADJUST=调整
	AdjustType string `json:"adjustType,omitempty" xml:"adjustType,omitempty"`
	// 总页数
	TotalPage int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
	// 当前页(从1开始)
	CurrentPage int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 每页记录的条数
	PageSize int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenInventoryReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolInventoryReportRequest = sync.Pool{
	New: func() any {
		return new(InventoryReportRequest)
	},
}

// GetInventoryReportRequest() 从对象池中获取InventoryReportRequest
func GetInventoryReportRequest() *InventoryReportRequest {
	return poolInventoryReportRequest.Get().(*InventoryReportRequest)
}

// ReleaseInventoryReportRequest 释放InventoryReportRequest
func ReleaseInventoryReportRequest(v *InventoryReportRequest) {
	v.Items = v.Items[:0]
	v.WarehouseCode = ""
	v.CheckOrderCode = ""
	v.CheckOrderId = ""
	v.OwnerCode = ""
	v.CheckTime = ""
	v.OutBizCode = ""
	v.Remark = ""
	v.AdjustType = ""
	v.TotalPage = 0
	v.CurrentPage = 0
	v.PageSize = 0
	v.ExtendProps = nil
	poolInventoryReportRequest.Put(v)
}
