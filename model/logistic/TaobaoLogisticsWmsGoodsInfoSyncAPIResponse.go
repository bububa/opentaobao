package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsGoodsInfoSyncAPIResponse WMS回传货品长宽高图片等信息 API返回值
// taobao.logistics.wms.goods.info.sync
//
// WMS回传货品长宽高图片等信息
type TaobaoLogisticsWmsGoodsInfoSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsGoodsInfoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel is WMS回传货品长宽高图片等信息 成功返回结果
type TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_goods_info_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误编码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否支持重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsGoodsInfoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Suc = false
	m.Retry = false
}

var poolTaobaoLogisticsWmsGoodsInfoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWmsGoodsInfoSyncAPIResponse)
	},
}

// GetTaobaoLogisticsWmsGoodsInfoSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWmsGoodsInfoSyncAPIResponse
func GetTaobaoLogisticsWmsGoodsInfoSyncAPIResponse() *TaobaoLogisticsWmsGoodsInfoSyncAPIResponse {
	return poolTaobaoLogisticsWmsGoodsInfoSyncAPIResponse.Get().(*TaobaoLogisticsWmsGoodsInfoSyncAPIResponse)
}

// ReleaseTaobaoLogisticsWmsGoodsInfoSyncAPIResponse 将 TaobaoLogisticsWmsGoodsInfoSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWmsGoodsInfoSyncAPIResponse(v *TaobaoLogisticsWmsGoodsInfoSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWmsGoodsInfoSyncAPIResponse.Put(v)
}
