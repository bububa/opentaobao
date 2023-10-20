package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagContentCancelAPIResponse 进口大包取消 API返回值
// cainiao.global.im.pickup.bigbag.content.cancel
//
// 进口大包取消
type CainiaoGlobalImPickupBigbagContentCancelAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagContentCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel is 进口大包取消 成功返回结果
type CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_content_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupBigbagContentCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupBigbagContentCancelAPIResponse)
	},
}

// GetCainiaoGlobalImPickupBigbagContentCancelAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagContentCancelAPIResponse
func GetCainiaoGlobalImPickupBigbagContentCancelAPIResponse() *CainiaoGlobalImPickupBigbagContentCancelAPIResponse {
	return poolCainiaoGlobalImPickupBigbagContentCancelAPIResponse.Get().(*CainiaoGlobalImPickupBigbagContentCancelAPIResponse)
}

// ReleaseCainiaoGlobalImPickupBigbagContentCancelAPIResponse 将 CainiaoGlobalImPickupBigbagContentCancelAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagContentCancelAPIResponse(v *CainiaoGlobalImPickupBigbagContentCancelAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagContentCancelAPIResponse.Put(v)
}
