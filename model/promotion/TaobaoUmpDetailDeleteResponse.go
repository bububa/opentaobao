package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除活动详情 APIResponse
taobao.ump.detail.delete

删除活动详情
*/
type TaobaoUmpDetailDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpDetailDeleteResponse `json:"taobao_ump_detail_delete_response,omitempty"`
}

type TaobaoUmpDetailDeleteResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
