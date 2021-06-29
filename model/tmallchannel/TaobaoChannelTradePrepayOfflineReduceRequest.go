package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（减少） API请求
taobao.channel.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少）
*/
type TaobaoChannelTradePrepayOfflineReduceRequest struct {
    model.Params
    // 减少流水
    _offlineReducePrepayParam   *TopOfflineReducePrepayDto
}

// 初始化TaobaoChannelTradePrepayOfflineReduceRequest对象
func NewTaobaoChannelTradePrepayOfflineReduceRequest() *TaobaoChannelTradePrepayOfflineReduceRequest{
    return &TaobaoChannelTradePrepayOfflineReduceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoChannelTradePrepayOfflineReduceRequest) GetApiMethodName() string {
    return "taobao.channel.trade.prepay.offline.reduce"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoChannelTradePrepayOfflineReduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoChannelTradePrepayOfflineReduceRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
    r._offlineReducePrepayParam = _offlineReducePrepayParam
    r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
    return nil
}

// OfflineReducePrepayParam Getter
func (r TaobaoChannelTradePrepayOfflineReduceRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
    return r._offlineReducePrepayParam
}
