package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询营销活动 APIResponse
taobao.ump.activity.get

查询营销活动
*/
type TaobaoUmpActivityGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpActivityGetResponse `json:"taobao_ump_activity_get_response,omitempty"`
}

type TaobaoUmpActivityGetResponse struct {

    // 营销活动的内容，可以通过ump sdk中的marketingTool来完成对该内容的处理
    Content   string `json:"content,omitempty"`

}
