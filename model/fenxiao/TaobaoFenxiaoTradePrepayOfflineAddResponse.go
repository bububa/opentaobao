package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下预存款流水增加 APIResponse
taobao.fenxiao.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加）
*/
type TaobaoFenxiaoTradePrepayOfflineAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoTradePrepayOfflineAddResponse `json:"taobao_fenxiao_trade_prepay_offline_add_response,omitempty"`
}

type TaobaoFenxiaoTradePrepayOfflineAddResponse struct {

    // result
    Result   *ResultTopDo `json:"result,omitempty"`

}
