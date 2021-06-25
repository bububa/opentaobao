package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情 APIResponse
taobao.ump.detail.get

查询活动详情
*/
type TaobaoUmpDetailGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpDetailGetResponse `json:"taobao_ump_detail_get_response,omitempty"`
}

type TaobaoUmpDetailGetResponse struct {

    // 活动详情信息，可以通过ump sdk中的MarketingTool来进行处理
    Content   string `json:"content,omitempty"`

}
