package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsSkuUpdateAPIResponse 商品信息的更新 API返回值
// taobao.wlb.wms.sku.update
//
// 商品信息的更新
type TaobaoWlbWmsSkuUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsSkuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsSkuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsSkuUpdateAPIResponseModel).Reset()
}

// TaobaoWlbWmsSkuUpdateAPIResponseModel is 商品信息的更新 成功返回结果
type TaobaoWlbWmsSkuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 错误编码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsSkuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WlErrorMsg = ""
	m.WlErrorCode = ""
	m.WlSuccess = false
}

var poolTaobaoWlbWmsSkuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsSkuUpdateAPIResponse)
	},
}

// GetTaobaoWlbWmsSkuUpdateAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsSkuUpdateAPIResponse
func GetTaobaoWlbWmsSkuUpdateAPIResponse() *TaobaoWlbWmsSkuUpdateAPIResponse {
	return poolTaobaoWlbWmsSkuUpdateAPIResponse.Get().(*TaobaoWlbWmsSkuUpdateAPIResponse)
}

// ReleaseTaobaoWlbWmsSkuUpdateAPIResponse 将 TaobaoWlbWmsSkuUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsSkuUpdateAPIResponse(v *TaobaoWlbWmsSkuUpdateAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsSkuUpdateAPIResponse.Put(v)
}
