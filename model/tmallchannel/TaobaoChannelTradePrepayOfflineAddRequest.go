package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（增加） APIRequest
taobao.channel.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加）
*/
type TaobaoChannelTradePrepayOfflineAddRequest struct {
    model.Params

    // 增加流水
    offlineAddPrepayParam   *TopOfflineAddPrepayDto 

}

func NewTaobaoChannelTradePrepayOfflineAddRequest() *TaobaoChannelTradePrepayOfflineAddRequest{
    return &TaobaoChannelTradePrepayOfflineAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoChannelTradePrepayOfflineAddRequest) GetApiMethodName() string {
    return "taobao.channel.trade.prepay.offline.add"
}

func (r TaobaoChannelTradePrepayOfflineAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoChannelTradePrepayOfflineAddRequest) SetOfflineAddPrepayParam(offlineAddPrepayParam *TopOfflineAddPrepayDto) error {
    r.offlineAddPrepayParam = offlineAddPrepayParam
    r.Set("offline_add_prepay_param", offlineAddPrepayParam)
    return nil
}

func (r TaobaoChannelTradePrepayOfflineAddRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDto {
    return r.offlineAddPrepayParam
}

