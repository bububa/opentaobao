package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
签收接口 APIResponse
tmall.msf.receive

签收接口
*/
type TmallMsfReceiveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_msf_receive_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   string `json:"result,omitempty" xml:"