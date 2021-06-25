package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（减少） APIResponse
taobao.fenxiao.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少）
*/
type TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoTradePrepayOfflineReduceResponse `json:"taobao_fenxiao_trade_prepay_offline_reduce_response,omitempty"`
}

type TaobaoFenxiaoTradePrepayOfflineReduceResponse struct {

    // result
    Result   *ResultTopDo `json:"result,omitempty"`

}
