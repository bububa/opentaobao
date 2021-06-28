package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户获取视频互动列表 APIResponse
taobao.interactive.list.getbyuser

根据用户来获取用户编辑的互动列表
*/
type TaobaoInteractiveListGetbyuserAPIResponse struct {
    model.CommonResponse
    TaobaoInteractiveListGetbyuserResponse
}

type TaobaoInteractiveListGetbyuserResponse struct {
    XMLName xml.Name `xml:"interactive_list_getbyuser_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoInteractiveListGetbyuserResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
