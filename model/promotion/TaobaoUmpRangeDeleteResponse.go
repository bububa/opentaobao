package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除活动范围 APIResponse
taobao.ump.range.delete

去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
*/
type TaobaoUmpRangeDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpRangeDeleteResponse `json:"taobao_ump_range_delete_response,omitempty"`
}

type TaobaoUmpRangeDeleteResponse struct {

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
