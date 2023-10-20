package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiCancelAPIResponse 商家取消获取的电子面单号 API返回值
// cainiao.waybill.ii.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaoWaybillIiCancelAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiCancelAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiCancelAPIResponseModel).Reset()
}

// CainiaoWaybillIiCancelAPIResponseModel is 商家取消获取的电子面单号 成功返回结果
type CainiaoWaybillIiCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用取消是否成功
	CancelResult bool `json:"cancel_result,omitempty" xml:"cancel_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CancelResult = false
}

var poolCainiaoWaybillIiCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiCancelAPIResponse)
	},
}

// GetCainiaoWaybillIiCancelAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiCancelAPIResponse
func GetCainiaoWaybillIiCancelAPIResponse() *CainiaoWaybillIiCancelAPIResponse {
	return poolCainiaoWaybillIiCancelAPIResponse.Get().(*CainiaoWaybillIiCancelAPIResponse)
}

// ReleaseCainiaoWaybillIiCancelAPIResponse 将 CainiaoWaybillIiCancelAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiCancelAPIResponse(v *CainiaoWaybillIiCancelAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiCancelAPIResponse.Put(v)
}
