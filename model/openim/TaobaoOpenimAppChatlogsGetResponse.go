package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
openim应用聊天记录查询 APIResponse
taobao.openim.app.chatlogs.get

查询openim应用的聊天记录
*/
type TaobaoOpenimAppChatlogsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimAppChatlogsGetResponse `json:"taobao_openim_app_chatlogs_get_response,omitempty"`
}

type TaobaoOpenimAppChatlogsGetResponse struct {

    // 查询结果
    Result   *EsMessageResult `json:"result,omitempty"`

}
