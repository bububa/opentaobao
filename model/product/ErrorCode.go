package product

import (
	"sync"
)

// ErrorCode 结构体
type ErrorCode struct {
	// 分组信息
	Fields []SeriesField `json:"fields,omitempty" xml:"fields>series_field,omitempty"`
	// 错误码
	MesCode string `json:"mes_code,omitempty" xml:"mes_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 市场
	Market string `json:"market,omitempty" xml:"market,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 系列名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 商品id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 用户id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 系列id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 系列扩展信息
	ItemSeriesExtend *ItemSeriesExtendDo `json:"item_series_extend,omitempty" xml:"item_series_extend,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 排序值
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 系列id
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
	// 特征
	Features int64 `json:"features,omitempty" xml:"features,omitempty"`
	// 分组id
	SegId int64 `json:"seg_id,omitempty" xml:"seg_id,omitempty"`
}

var poolErrorCode = sync.Pool{
	New: func() any {
		return new(ErrorCode)
	},
}

// GetErrorCode() 从对象池中获取ErrorCode
func GetErrorCode() *ErrorCode {
	return poolErrorCode.Get().(*ErrorCode)
}

// ReleaseErrorCode 释放ErrorCode
func ReleaseErrorCode(v *ErrorCode) {
	v.Fields = v.Fields[:0]
	v.MesCode = ""
	v.Message = ""
	v.Market = ""
	v.GmtModified = ""
	v.Name = ""
	v.GmtCreate = ""
	v.OuterId = ""
	v.SellerId = 0
	v.Id = 0
	v.ItemSeriesExtend = nil
	v.Status = 0
	v.Sort = 0
	v.SeriesId = 0
	v.Features = 0
	v.SegId = 0
	poolErrorCode.Put(v)
}
