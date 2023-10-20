package vaccin

import (
	"sync"
)

// PovSubscribeDetailModel 结构体
type PovSubscribeDetailModel struct {
	// 预约者姓名
	SubscriberName string `json:"subscriber_name,omitempty" xml:"subscriber_name,omitempty"`
	// 预约者电话
	SubscriberMobile string `json:"subscriber_mobile,omitempty" xml:"subscriber_mobile,omitempty"`
	// 身份证号
	IdCardMd5 string `json:"id_card_md5,omitempty" xml:"id_card_md5,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 预约时间
	SubscribeTime string `json:"subscribe_time,omitempty" xml:"subscribe_time,omitempty"`
	// 取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// pov名称
	PovName string `json:"pov_name,omitempty" xml:"pov_name,omitempty"`
	// isvorderId
	IsvOrderId string `json:"isv_order_id,omitempty" xml:"isv_order_id,omitempty"`
	// povId
	PovId int64 `json:"pov_id,omitempty" xml:"pov_id,omitempty"`
	// 省code
	PovProvinceCode int64 `json:"pov_province_code,omitempty" xml:"pov_province_code,omitempty"`
	// 市code
	PovCityCode int64 `json:"pov_city_code,omitempty" xml:"pov_city_code,omitempty"`
	// 区code
	AreaCode int64 `json:"area_code,omitempty" xml:"area_code,omitempty"`
}

var poolPovSubscribeDetailModel = sync.Pool{
	New: func() any {
		return new(PovSubscribeDetailModel)
	},
}

// GetPovSubscribeDetailModel() 从对象池中获取PovSubscribeDetailModel
func GetPovSubscribeDetailModel() *PovSubscribeDetailModel {
	return poolPovSubscribeDetailModel.Get().(*PovSubscribeDetailModel)
}

// ReleasePovSubscribeDetailModel 释放PovSubscribeDetailModel
func ReleasePovSubscribeDetailModel(v *PovSubscribeDetailModel) {
	v.SubscriberName = ""
	v.SubscriberMobile = ""
	v.IdCardMd5 = ""
	v.GmtCreate = ""
	v.SubscribeTime = ""
	v.CancelTime = ""
	v.PovName = ""
	v.IsvOrderId = ""
	v.PovId = 0
	v.PovProvinceCode = 0
	v.PovCityCode = 0
	v.AreaCode = 0
	poolPovSubscribeDetailModel.Put(v)
}
