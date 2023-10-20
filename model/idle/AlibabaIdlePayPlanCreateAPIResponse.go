package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdlePayPlanCreateAPIResponse 创建代扣计划 API返回值
// alibaba.idle.pay.plan.create
//
// 闲鱼平台代扣能力：
// 1、用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
// 2、创建代扣计划
type AlibabaIdlePayPlanCreateAPIResponse struct {
	model.CommonResponse
	AlibabaIdlePayPlanCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdlePayPlanCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdlePayPlanCreateAPIResponseModel).Reset()
}

// AlibabaIdlePayPlanCreateAPIResponseModel is 创建代扣计划 成功返回结果
type AlibabaIdlePayPlanCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_pay_plan_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdlePayPlanCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdlePayPlanCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdlePayPlanCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdlePayPlanCreateAPIResponse)
	},
}

// GetAlibabaIdlePayPlanCreateAPIResponse 从 sync.Pool 获取 AlibabaIdlePayPlanCreateAPIResponse
func GetAlibabaIdlePayPlanCreateAPIResponse() *AlibabaIdlePayPlanCreateAPIResponse {
	return poolAlibabaIdlePayPlanCreateAPIResponse.Get().(*AlibabaIdlePayPlanCreateAPIResponse)
}

// ReleaseAlibabaIdlePayPlanCreateAPIResponse 将 AlibabaIdlePayPlanCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdlePayPlanCreateAPIResponse(v *AlibabaIdlePayPlanCreateAPIResponse) {
	v.Reset()
	poolAlibabaIdlePayPlanCreateAPIResponse.Put(v)
}
