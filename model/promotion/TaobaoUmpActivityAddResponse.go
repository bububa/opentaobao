package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增优惠活动 APIResponse
taobao.ump.activity.add

新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
*/
type TaobaoUmpActivityAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpActivityAddResponse `json:"taobao_ump_activity_add_response,omitempty"`
}

type TaobaoUmpActivityAddResponse struct {

    // 活动id
    ActId   int64 `json:"act_id,omitempty"`

}
