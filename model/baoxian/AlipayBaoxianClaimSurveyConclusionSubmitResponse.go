package baoxian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保险退货服务商勘察结论提交接口 APIResponse
alipay.baoxian.claim.survey.conclusion.submit

保险退货服务商提交勘察结论
*/
type AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alipay_baoxian_claim_survey_conclusion_submit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *AliSceneResult `json:"result,omitempty" xml:"