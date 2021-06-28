package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
集货鉴定结果 APIResponse
taobao.wlb.imports.vas.identity.result

集货鉴定结果查询
*/
type TaobaoWlbImportsVasIdentityResultAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_imports_vas_identity_result_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回出参数结果
    
    Result   *TopResult `json:"result,omitempty" xml:"