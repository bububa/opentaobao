package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNextoneLogisticsSignUpdateAPIResponse AG物流签收状态写接口 API返回值
// taobao.nextone.logistics.sign.update
//
// 商家上传退货的签收状态给AG
type TaobaoNextoneLogisticsSignUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoNextoneLogisticsSignUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoNextoneLogisticsSignUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoNextoneLogisticsSignUpdateAPIResponseModel).Reset()
}

// TaobaoNextoneLogisticsSignUpdateAPIResponseModel is AG物流签收状态写接口 成功返回结果
type TaobaoNextoneLogisticsSignUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"nextone_logistics_sign_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoNextoneLogisticsSignUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoNextoneLogisticsSignUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoNextoneLogisticsSignUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoNextoneLogisticsSignUpdateAPIResponse)
	},
}

// GetTaobaoNextoneLogisticsSignUpdateAPIResponse 从 sync.Pool 获取 TaobaoNextoneLogisticsSignUpdateAPIResponse
func GetTaobaoNextoneLogisticsSignUpdateAPIResponse() *TaobaoNextoneLogisticsSignUpdateAPIResponse {
	return poolTaobaoNextoneLogisticsSignUpdateAPIResponse.Get().(*TaobaoNextoneLogisticsSignUpdateAPIResponse)
}

// ReleaseTaobaoNextoneLogisticsSignUpdateAPIResponse 将 TaobaoNextoneLogisticsSignUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoNextoneLogisticsSignUpdateAPIResponse(v *TaobaoNextoneLogisticsSignUpdateAPIResponse) {
	v.Reset()
	poolTaobaoNextoneLogisticsSignUpdateAPIResponse.Put(v)
}
