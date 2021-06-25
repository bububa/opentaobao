package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅核销接口 APIResponse
tmall.msf.verify

msf服务核销的top接口
*/
type TmallMsfVerifyAPIResponse struct {
    model.CommonResponse
    Response *TmallMsfVerifyResponse `json:"tmall_msf_verify_response,omitempty"`
}

type TmallMsfVerifyResponse struct {

    // result
    Result   string `json:"result,omitempty"`

}
