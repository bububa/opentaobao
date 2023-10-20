package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverParcelQueryAPIResponse 获取交接单小包信息 API返回值
// cainiao.global.handover.parcel.query
//
// 提供给ISV通过该接口查询小包信息
type CainiaoGlobalHandoverParcelQueryAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverParcelQueryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverParcelQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalHandoverParcelQueryAPIResponseModel).Reset()
}

// CainiaoGlobalHandoverParcelQueryAPIResponseModel is 获取交接单小包信息 成功返回结果
type CainiaoGlobalHandoverParcelQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_parcel_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalHandoverParcelQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalHandoverParcelQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalHandoverParcelQueryAPIResponse)
	},
}

// GetCainiaoGlobalHandoverParcelQueryAPIResponse 从 sync.Pool 获取 CainiaoGlobalHandoverParcelQueryAPIResponse
func GetCainiaoGlobalHandoverParcelQueryAPIResponse() *CainiaoGlobalHandoverParcelQueryAPIResponse {
	return poolCainiaoGlobalHandoverParcelQueryAPIResponse.Get().(*CainiaoGlobalHandoverParcelQueryAPIResponse)
}

// ReleaseCainiaoGlobalHandoverParcelQueryAPIResponse 将 CainiaoGlobalHandoverParcelQueryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalHandoverParcelQueryAPIResponse(v *CainiaoGlobalHandoverParcelQueryAPIResponse) {
	v.Reset()
	poolCainiaoGlobalHandoverParcelQueryAPIResponse.Put(v)
}
