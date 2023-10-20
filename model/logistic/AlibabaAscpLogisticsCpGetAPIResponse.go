package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsCpGetAPIResponse 快递公司资源列表查询接口 API返回值
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
type AlibabaAscpLogisticsCpGetAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsCpGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsCpGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsCpGetAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsCpGetAPIResponseModel is 快递公司资源列表查询接口 成功返回结果
type AlibabaAscpLogisticsCpGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_cp_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsCpGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsCpGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsCpGetAPIResponse)
	},
}

// GetAlibabaAscpLogisticsCpGetAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsCpGetAPIResponse
func GetAlibabaAscpLogisticsCpGetAPIResponse() *AlibabaAscpLogisticsCpGetAPIResponse {
	return poolAlibabaAscpLogisticsCpGetAPIResponse.Get().(*AlibabaAscpLogisticsCpGetAPIResponse)
}

// ReleaseAlibabaAscpLogisticsCpGetAPIResponse 将 AlibabaAscpLogisticsCpGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsCpGetAPIResponse(v *AlibabaAscpLogisticsCpGetAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsCpGetAPIResponse.Put(v)
}
