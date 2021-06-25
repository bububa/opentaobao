package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
口碑综合体订单列表信息查询 APIResponse
taobao.koubei.tribe.open.order.page

查询口碑商圈用户的订单列表信息
*/
type TaobaoKoubeiTribeOpenOrderPageAPIResponse struct {
    model.CommonResponse
    Response *TaobaoKoubeiTribeOpenOrderPageResponse `json:"taobao_koubei_tribe_open_order_page_response,omitempty"`
}

type TaobaoKoubeiTribeOpenOrderPageResponse struct {

    // 接口返回model
    Result   *TaobaoKoubeiTribeOpenOrderPageResult `json:"result,omitempty"`

}
