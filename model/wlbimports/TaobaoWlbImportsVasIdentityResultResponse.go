package wlbimports

import (
    "github.com/bububa/opentaobao/model"
)

/* 
集货鉴定结果 APIResponse
taobao.wlb.imports.vas.identity.result

集货鉴定结果查询
*/
type TaobaoWlbImportsVasIdentityResultAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbImportsVasIdentityResultResponse `json:"wlb_imports_vas_identity_result_response,omitempty"` 
    TaobaoWlbImportsVasIdentityResultResponse
}

/* model for simplify = false
type TaobaoWlbImportsVasIdentityResultResponse struct {

    // 返回出参数结果
    
    Result  *struct {
        TopResult  *TopResult `json:"top_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWlbImportsVasIdentityResultResponse struct {

    // 返回出参数结果
    Result   *TopResult `json:"result,omitempty"`

}
