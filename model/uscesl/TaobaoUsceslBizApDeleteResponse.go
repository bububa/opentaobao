package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除价签AP设备 APIResponse
taobao.uscesl.biz.ap.delete

删除价签AP设备
*/
type TaobaoUsceslBizApDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizApDeleteResponse
}

type TaobaoUsceslBizApDeleteResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_ap_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功与否看result.success，返回true或者false
    
    Result   *TaobaoUsceslBizApDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
