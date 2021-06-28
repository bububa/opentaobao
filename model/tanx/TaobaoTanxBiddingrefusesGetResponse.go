package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tanx竞价失败反馈api APIResponse
taobao.tanx.biddingrefuses.get

竞价失败反馈根据创意id查询API提供
*/
type TaobaoTanxBiddingrefusesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_biddingrefuses_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"