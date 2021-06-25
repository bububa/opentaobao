package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（减少） APIRequest
taobao.fenxiao.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少）
*/
type TaobaoFenxiaoTradePrepayOfflineReduceRequest struct {
    model.Params

    // 减少流水
    offlineReducePrepayParam   *TopOfflineReducePrepayDto 

}

func NewTaobaoFenxiaoTradePrepayOfflineReduceRequest() *TaobaoFenxiaoTradePrepayOfflineReduceRequest{
    return &TaobaoFenxiaoTradePrepayOfflineReduceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetApiMethodName() string {
    return "taobao.fenxiao.trade.prepay.offline.reduce"
}

func (r TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoTradePrepayOfflineReduceRequest) SetOfflineReducePrepayParam(offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
    r.offlineReducePrepayParam = offlineReducePrepayParam
    r.Set("offline_reduce_prepay_param", offlineReducePrepayParam)
    return nil
}

func (r TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
    return r.offlineReducePrepayParam
}

