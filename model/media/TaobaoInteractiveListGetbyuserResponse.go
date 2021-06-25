package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户获取视频互动列表 APIResponse
taobao.interactive.list.getbyuser

根据用户来获取用户编辑的互动列表
*/
type TaobaoInteractiveListGetbyuserAPIResponse struct {
    model.CommonResponse
    Response *TaobaoInteractiveListGetbyuserResponse `json:"taobao_interactive_list_getbyuser_response,omitempty"`
}

type TaobaoInteractiveListGetbyuserResponse struct {

    // result
    Result   *TaobaoInteractiveListGetbyuserResult `json:"result,omitempty"`

}
