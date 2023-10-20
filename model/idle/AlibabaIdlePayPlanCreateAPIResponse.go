package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlepayplancreateAPIResponse 创建代扣计划 API返回值
// alibaba.idle.pay.plan.create
//
// 闲鱼平台代扣能力：
// 1、用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
// 2、创建代扣计划
type AlibabaidlepayplancreateAPIResponse struct {
	model.CommonResponse
	AlibabaidlepayplancreateAPIResponseModel
}

// AlibabaidlepayplancreateAPIResponseModel is 创建代扣计划 成功返回结果
type AlibabaidlepayplancreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_pay_plan_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaidlepayplancreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
