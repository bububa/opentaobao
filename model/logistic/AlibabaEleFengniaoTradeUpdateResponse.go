package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新蜂鸟扣费状态 APIResponse
alibaba.ele.fengniao.trade.update

汇金扣费成功后，回调该接口更新扣费状态
*/
type AlibabaEleFengniaoTradeUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoTradeUpdateResponse `json:"alibaba_ele_fengniao_trade_update_response,omitempty"`
}

type AlibabaEleFengniaoTradeUpdateResponse struct {

    // 1:成功 0：失败
    Status   int64 `json:"status,omitempty"`

    // 无此交易记录
    ErrorMsg   string `json:"error_msg,omitempty"`

}
