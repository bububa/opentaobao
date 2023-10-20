package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsSkuCreateAPIResponse 商品同步 API返回值
// taobao.wlb.wms.sku.create
//
// 商品同步
type TaobaoWlbWmsSkuCreateAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsSkuCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsSkuCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsSkuCreateAPIResponseModel).Reset()
}

// TaobaoWlbWmsSkuCreateAPIResponseModel is 商品同步 成功返回结果
type TaobaoWlbWmsSkuCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_sku_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 错误码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 系统自动生成
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsSkuCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WlErrorMsg = ""
	m.WlErrorCode = ""
	m.ItemId = 0
	m.WlSuccess = false
}

var poolTaobaoWlbWmsSkuCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsSkuCreateAPIResponse)
	},
}

// GetTaobaoWlbWmsSkuCreateAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsSkuCreateAPIResponse
func GetTaobaoWlbWmsSkuCreateAPIResponse() *TaobaoWlbWmsSkuCreateAPIResponse {
	return poolTaobaoWlbWmsSkuCreateAPIResponse.Get().(*TaobaoWlbWmsSkuCreateAPIResponse)
}

// ReleaseTaobaoWlbWmsSkuCreateAPIResponse 将 TaobaoWlbWmsSkuCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsSkuCreateAPIResponse(v *TaobaoWlbWmsSkuCreateAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsSkuCreateAPIResponse.Put(v)
}
