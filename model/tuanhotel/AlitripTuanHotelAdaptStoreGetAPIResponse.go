package tuanhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelAdaptStoreGetAPIResponse 酒店团购套餐关联适用门店 API返回值
// alitrip.tuan.hotel.adapt.store.get
//
// 输入shid，返回关联门店详情信息
type AlitripTuanHotelAdaptStoreGetAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelAdaptStoreGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTuanHotelAdaptStoreGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTuanHotelAdaptStoreGetAPIResponseModel).Reset()
}

// AlitripTuanHotelAdaptStoreGetAPIResponseModel is 酒店团购套餐关联适用门店 成功返回结果
type AlitripTuanHotelAdaptStoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_adapt_store_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关联门店列表
	StoreDetailList []StoreDetailVoList `json:"store_detail_list,omitempty" xml:"store_detail_list>store_detail_vo_list,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTuanHotelAdaptStoreGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StoreDetailList = m.StoreDetailList[:0]
	m.MsgCode = ""
	m.Message = ""
	m.Status = false
}

var poolAlitripTuanHotelAdaptStoreGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTuanHotelAdaptStoreGetAPIResponse)
	},
}

// GetAlitripTuanHotelAdaptStoreGetAPIResponse 从 sync.Pool 获取 AlitripTuanHotelAdaptStoreGetAPIResponse
func GetAlitripTuanHotelAdaptStoreGetAPIResponse() *AlitripTuanHotelAdaptStoreGetAPIResponse {
	return poolAlitripTuanHotelAdaptStoreGetAPIResponse.Get().(*AlitripTuanHotelAdaptStoreGetAPIResponse)
}

// ReleaseAlitripTuanHotelAdaptStoreGetAPIResponse 将 AlitripTuanHotelAdaptStoreGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripTuanHotelAdaptStoreGetAPIResponse(v *AlitripTuanHotelAdaptStoreGetAPIResponse) {
	v.Reset()
	poolAlitripTuanHotelAdaptStoreGetAPIResponse.Put(v)
}
