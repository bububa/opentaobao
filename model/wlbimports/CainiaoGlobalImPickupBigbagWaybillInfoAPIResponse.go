package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse 大包面单查询 API返回值
// cainiao.global.im.pickup.bigbag.waybill.info
//
// 大包面单查询
type CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagWaybillInfoAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupBigbagWaybillInfoAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupBigbagWaybillInfoAPIResponseModel is 大包面单查询 成功返回结果
type CainiaoGlobalImPickupBigbagWaybillInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_waybill_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagWaybillInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupBigbagWaybillInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse)
	},
}

// GetCainiaoGlobalImPickupBigbagWaybillInfoAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse
func GetCainiaoGlobalImPickupBigbagWaybillInfoAPIResponse() *CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse {
	return poolCainiaoGlobalImPickupBigbagWaybillInfoAPIResponse.Get().(*CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse)
}

// ReleaseCainiaoGlobalImPickupBigbagWaybillInfoAPIResponse 将 CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagWaybillInfoAPIResponse(v *CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagWaybillInfoAPIResponse.Put(v)
}
