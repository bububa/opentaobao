package baoxian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse 保险退货服务商勘察结论提交接口 API返回值
// alipay.baoxian.claim.survey.conclusion.submit
//
// 保险退货服务商提交勘察结论
type AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse struct {
	model.CommonResponse
	AlipayBaoxianClaimSurveyConclusionSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlipayBaoxianClaimSurveyConclusionSubmitAPIResponseModel).Reset()
}

// AlipayBaoxianClaimSurveyConclusionSubmitAPIResponseModel is 保险退货服务商勘察结论提交接口 成功返回结果
type AlipayBaoxianClaimSurveyConclusionSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_baoxian_claim_survey_conclusion_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AliSceneResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimSurveyConclusionSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlipayBaoxianClaimSurveyConclusionSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse)
	},
}

// GetAlipayBaoxianClaimSurveyConclusionSubmitAPIResponse 从 sync.Pool 获取 AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse
func GetAlipayBaoxianClaimSurveyConclusionSubmitAPIResponse() *AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse {
	return poolAlipayBaoxianClaimSurveyConclusionSubmitAPIResponse.Get().(*AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse)
}

// ReleaseAlipayBaoxianClaimSurveyConclusionSubmitAPIResponse 将 AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlipayBaoxianClaimSurveyConclusionSubmitAPIResponse(v *AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse) {
	v.Reset()
	poolAlipayBaoxianClaimSurveyConclusionSubmitAPIResponse.Put(v)
}
