package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiQueryByWaybillcodeAPIResponse 通过面单号查询面单打印报文 API返回值
// cainiao.waybill.ii.query.by.waybillcode
//
// 通过面单号查询面单的打印报文
type CainiaoWaybillIiQueryByWaybillcodeAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiQueryByWaybillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiQueryByWaybillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiQueryByWaybillcodeAPIResponseModel).Reset()
}

// CainiaoWaybillIiQueryByWaybillcodeAPIResponseModel is 通过面单号查询面单打印报文 成功返回结果
type CainiaoWaybillIiQueryByWaybillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_query_by_waybillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回值
	Modules []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_with_result_desc_response,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiQueryByWaybillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Modules = m.Modules[:0]
}

var poolCainiaoWaybillIiQueryByWaybillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiQueryByWaybillcodeAPIResponse)
	},
}

// GetCainiaoWaybillIiQueryByWaybillcodeAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiQueryByWaybillcodeAPIResponse
func GetCainiaoWaybillIiQueryByWaybillcodeAPIResponse() *CainiaoWaybillIiQueryByWaybillcodeAPIResponse {
	return poolCainiaoWaybillIiQueryByWaybillcodeAPIResponse.Get().(*CainiaoWaybillIiQueryByWaybillcodeAPIResponse)
}

// ReleaseCainiaoWaybillIiQueryByWaybillcodeAPIResponse 将 CainiaoWaybillIiQueryByWaybillcodeAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiQueryByWaybillcodeAPIResponse(v *CainiaoWaybillIiQueryByWaybillcodeAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiQueryByWaybillcodeAPIResponse.Put(v)
}
