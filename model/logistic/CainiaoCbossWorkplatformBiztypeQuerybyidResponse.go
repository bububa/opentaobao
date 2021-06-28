package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据业务类型id查询业务类型详细信息 APIResponse
cainiao.cboss.workplatform.biztype.querybyid

菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCbossWorkplatformBiztypeQuerybyidResponse `json:"cainiao_cboss_workplatform_biztype_querybyid_response,omitempty"` 
    CainiaoCbossWorkplatformBiztypeQuerybyidResponse
}

/* model for simplify = false
type CainiaoCbossWorkplatformBiztypeQuerybyidResponse struct {

    // result
    
    Result  *struct {
        CainiaoCbossWorkplatformBiztypeQuerybyidResult  *CainiaoCbossWorkplatformBiztypeQuerybyidResult `json:"cainiao_cboss_workplatform_biztype_querybyid_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoCbossWorkplatformBiztypeQuerybyidResponse struct {

    // result
    Result   *CainiaoCbossWorkplatformBiztypeQuerybyidResult `json:"result,omitempty"`

}
