package baoxian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
保险退货服务商勘察结论提交接口 APIResponse
alipay.baoxian.claim.survey.conclusion.submit

保险退货服务商提交勘察结论
*/
type AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse struct {
    model.CommonResponse
    // Response *AlipayBaoxianClaimSurveyConclusionSubmitResponse `json:"alipay_baoxian_claim_survey_conclusion_submit_response,omitempty"` 
    AlipayBaoxianClaimSurveyConclusionSubmitResponse
}

/* model for simplify = false
type AlipayBaoxianClaimSurveyConclusionSubmitResponse struct {

    // 系统自动生成
    
    Result  *struct {
        AliSceneResult  *AliSceneResult `json:"ali_scene_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlipayBaoxianClaimSurveyConclusionSubmitResponse struct {

    // 系统自动生成
    Result   *AliSceneResult `json:"result,omitempty"`

}
