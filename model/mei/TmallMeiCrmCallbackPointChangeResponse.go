package mei

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌积分变更回调API APIResponse
tmall.mei.crm.callback.point.change

线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
*/
type TmallMeiCrmCallbackPointChangeAPIResponse struct {
    model.CommonResponse
    TmallMeiCrmCallbackPointChangeResponse
}

type TmallMeiCrmCallbackPointChangeResponse struct {
    XMLName xml.Name `xml:"tmall_mei_crm_callback_point_change_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
}
