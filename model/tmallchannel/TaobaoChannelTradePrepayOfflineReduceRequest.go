package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（减少） APIRequest
taobao.channel.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少）
*/
type TaobaoChannelTradePrepayOfflineReduceRequest struct {
    model.Params

    // 减少流水
    offlineReducePrepayParam   *TopOfflineReducePrepayDto 

}

func NewTaobaoChannelTradePrepayOfflineReduceRequest() *TaobaoChannelTradePrepayOfflineReduceRequest{
    return &TaobaoChannelTradePrepayOfflineReduceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoChannelTradePrepayOfflineReduceRequest) GetApiMethodName() string {
    return "taobao.channel.trade.prepay.offline.reduce"
}

func (r TaobaoChannelTradePrepayOfflineReduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoChannelTradePrepayOfflineReduceRequest) SetOfflineReducePrepayParam(offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
    r.offlineReducePrepayParam = offlineReducePrepayParam
    r.Set("offline_reduce_prepay_param", offlineReducePrepayParam)
    return nil
}

func (r TaobaoChannelTradePrepayOfflineReduceRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
    return r.offlineReducePrepayParam
}

