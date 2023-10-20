package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagContentCreateAPIResponse 大包创建 API返回值
// cainiao.global.im.pickup.bigbag.content.create
//
// 大包创建
type CainiaoGlobalImPickupBigbagContentCreateAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagContentCreateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagContentCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupBigbagContentCreateAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupBigbagContentCreateAPIResponseModel is 大包创建 成功返回结果
type CainiaoGlobalImPickupBigbagContentCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_content_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagContentCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupBigbagContentCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupBigbagContentCreateAPIResponse)
	},
}

// GetCainiaoGlobalImPickupBigbagContentCreateAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagContentCreateAPIResponse
func GetCainiaoGlobalImPickupBigbagContentCreateAPIResponse() *CainiaoGlobalImPickupBigbagContentCreateAPIResponse {
	return poolCainiaoGlobalImPickupBigbagContentCreateAPIResponse.Get().(*CainiaoGlobalImPickupBigbagContentCreateAPIResponse)
}

// ReleaseCainiaoGlobalImPickupBigbagContentCreateAPIResponse 将 CainiaoGlobalImPickupBigbagContentCreateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagContentCreateAPIResponse(v *CainiaoGlobalImPickupBigbagContentCreateAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagContentCreateAPIResponse.Put(v)
}
