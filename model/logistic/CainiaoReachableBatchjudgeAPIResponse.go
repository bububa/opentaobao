package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoReachableBatchjudgeAPIResponse 是否派送可达判定批量查询接口 API返回值
// cainiao.reachable.batchjudge
//
// 提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型
type CainiaoReachableBatchjudgeAPIResponse struct {
	model.CommonResponse
	CainiaoReachableBatchjudgeAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoReachableBatchjudgeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoReachableBatchjudgeAPIResponseModel).Reset()
}

// CainiaoReachableBatchjudgeAPIResponseModel is 是否派送可达判定批量查询接口 成功返回结果
type CainiaoReachableBatchjudgeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_reachable_batchjudge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoReachableBatchjudgeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoReachableBatchjudgeAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoReachableBatchjudgeAPIResponse)
	},
}

// GetCainiaoReachableBatchjudgeAPIResponse 从 sync.Pool 获取 CainiaoReachableBatchjudgeAPIResponse
func GetCainiaoReachableBatchjudgeAPIResponse() *CainiaoReachableBatchjudgeAPIResponse {
	return poolCainiaoReachableBatchjudgeAPIResponse.Get().(*CainiaoReachableBatchjudgeAPIResponse)
}

// ReleaseCainiaoReachableBatchjudgeAPIResponse 将 CainiaoReachableBatchjudgeAPIResponse 保存到 sync.Pool
func ReleaseCainiaoReachableBatchjudgeAPIResponse(v *CainiaoReachableBatchjudgeAPIResponse) {
	v.Reset()
	poolCainiaoReachableBatchjudgeAPIResponse.Put(v)
}
