package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改活动信息 APIResponse
taobao.ump.activity.update

修改营销活动
*/
type TaobaoUmpActivityUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpActivityUpdateResponse `json:"taobao_ump_activity_update_response,omitempty"`
}

type TaobaoUmpActivityUpdateResponse struct {

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
