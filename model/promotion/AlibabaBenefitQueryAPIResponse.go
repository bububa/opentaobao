package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBenefitQueryAPIResponse 奖池奖品查询列表 API返回值
// alibaba.benefit.query
//
// 功能：奖池奖品查询列表
// 业务逻辑：程序中通过奖池编号ename,业务方身份appName来查询奖池提供的奖品返回给
// 小程。
// 安全保障：为保证数据不会越权，需要卖家授，并且验证系统参数appKey。只有通过授权的，并且
// appkey验证通过的，才会查数据 并透出，否则直接失败。
// 因为appkey是系统参数，并且程序内部可以验证appkey和业务身份appName的关系
// 是否一致，所以可以保证参数appName的合法性，没有越权。
type AlibabaBenefitQueryAPIResponse struct {
	model.CommonResponse
	AlibabaBenefitQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBenefitQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBenefitQueryAPIResponseModel).Reset()
}

// AlibabaBenefitQueryAPIResponseModel is 奖池奖品查询列表 成功返回结果
type AlibabaBenefitQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_benefit_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaBenefitQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBenefitQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBenefitQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBenefitQueryAPIResponse)
	},
}

// GetAlibabaBenefitQueryAPIResponse 从 sync.Pool 获取 AlibabaBenefitQueryAPIResponse
func GetAlibabaBenefitQueryAPIResponse() *AlibabaBenefitQueryAPIResponse {
	return poolAlibabaBenefitQueryAPIResponse.Get().(*AlibabaBenefitQueryAPIResponse)
}

// ReleaseAlibabaBenefitQueryAPIResponse 将 AlibabaBenefitQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBenefitQueryAPIResponse(v *AlibabaBenefitQueryAPIResponse) {
	v.Reset()
	poolAlibabaBenefitQueryAPIResponse.Put(v)
}
