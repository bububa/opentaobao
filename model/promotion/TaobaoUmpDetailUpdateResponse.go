package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改活动详情 APIResponse
taobao.ump.detail.update

更新活动详情
*/
type TaobaoUmpDetailUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpDetailUpdateResponse `json:"taobao_ump_detail_update_response,omitempty"`
}

type TaobaoUmpDetailUpdateResponse struct {

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
