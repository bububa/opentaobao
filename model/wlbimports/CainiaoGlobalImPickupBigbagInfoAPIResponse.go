package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagInfoAPIResponse 大包状态查询 API返回值
// cainiao.global.im.pickup.bigbag.info
//
// 大包状态查询
type CainiaoGlobalImPickupBigbagInfoAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagInfoAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupBigbagInfoAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupBigbagInfoAPIResponseModel is 大包状态查询 成功返回结果
type CainiaoGlobalImPickupBigbagInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupBigbagInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupBigbagInfoAPIResponse)
	},
}

// GetCainiaoGlobalImPickupBigbagInfoAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagInfoAPIResponse
func GetCainiaoGlobalImPickupBigbagInfoAPIResponse() *CainiaoGlobalImPickupBigbagInfoAPIResponse {
	return poolCainiaoGlobalImPickupBigbagInfoAPIResponse.Get().(*CainiaoGlobalImPickupBigbagInfoAPIResponse)
}

// ReleaseCainiaoGlobalImPickupBigbagInfoAPIResponse 将 CainiaoGlobalImPickupBigbagInfoAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagInfoAPIResponse(v *CainiaoGlobalImPickupBigbagInfoAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagInfoAPIResponse.Put(v)
}
