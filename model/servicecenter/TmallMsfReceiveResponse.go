package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
签收接口 APIResponse
tmall.msf.receive

签收接口
*/
type TmallMsfReceiveAPIResponse struct {
    model.CommonResponse
    Response *TmallMsfReceiveResponse `json:"tmall_msf_receive_response,omitempty"`
}

type TmallMsfReceiveResponse struct {

    // result
    Result   string `json:"result,omitempty"`

}
