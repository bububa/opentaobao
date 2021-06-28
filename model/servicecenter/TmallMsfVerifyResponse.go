package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅核销接口 APIResponse
tmall.msf.verify

msf服务核销的top接口
*/
type TmallMsfVerifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_msf_verify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   string `json:"result,omitempty" xml:"