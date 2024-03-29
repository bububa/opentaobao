package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiQueryByTradecodeAPIResponse 通过订单号查询电子面单通接口 API返回值
// cainiao.waybill.ii.query.by.tradecode
//
// 通过订单号查看面单的信息
type CainiaoWaybillIiQueryByTradecodeAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiQueryByTradecodeAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiQueryByTradecodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiQueryByTradecodeAPIResponseModel).Reset()
}

// CainiaoWaybillIiQueryByTradecodeAPIResponseModel is 通过订单号查询电子面单通接口 成功返回结果
type CainiaoWaybillIiQueryByTradecodeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_query_by_tradecode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回值
	Modules []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_with_result_desc_response,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiQueryByTradecodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Modules = m.Modules[:0]
}

var poolCainiaoWaybillIiQueryByTradecodeAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiQueryByTradecodeAPIResponse)
	},
}

// GetCainiaoWaybillIiQueryByTradecodeAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiQueryByTradecodeAPIResponse
func GetCainiaoWaybillIiQueryByTradecodeAPIResponse() *CainiaoWaybillIiQueryByTradecodeAPIResponse {
	return poolCainiaoWaybillIiQueryByTradecodeAPIResponse.Get().(*CainiaoWaybillIiQueryByTradecodeAPIResponse)
}

// ReleaseCainiaoWaybillIiQueryByTradecodeAPIResponse 将 CainiaoWaybillIiQueryByTradecodeAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiQueryByTradecodeAPIResponse(v *CainiaoWaybillIiQueryByTradecodeAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiQueryByTradecodeAPIResponse.Put(v)
}
