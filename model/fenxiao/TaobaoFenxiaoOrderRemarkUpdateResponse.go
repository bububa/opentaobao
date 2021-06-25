package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改采购单备注 APIResponse
taobao.fenxiao.order.remark.update

供应商修改采购单备注
*/
type TaobaoFenxiaoOrderRemarkUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoOrderRemarkUpdateResponse `json:"taobao_fenxiao_order_remark_update_response,omitempty"`
}

type TaobaoFenxiaoOrderRemarkUpdateResponse struct {

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
