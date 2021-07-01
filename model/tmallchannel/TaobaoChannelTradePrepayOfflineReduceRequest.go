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
type TaobaoChannelTradePrepayOfflineReduceAPIRequest struct {
    model.Params
    // 减少流水
    _offlineReducePrepayParam   *TopOfflineReducePrepayDTO
}

// 初始化TaobaoChannelTradePrepayOfflineReduceAPIRequest对象
func NewTaobaoChannelTradePrepayOfflineReduceRequest() *TaobaoChannelTradePrepayOfflineReduceAPIRequest{
    return &TaobaoChannelTradePrepayOfflineReduceAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetApiMethodName() string {
    return "taobao.channel.trade.prepay.offline.reduce"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoChannelTradePrepayOfflineReduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDTO) error {
    r._offlineReducePrepayParam = _offlineReducePrepayParam
    r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
    return nil
}

// OfflineReducePrepayParam Getter
func (r TaobaoChannelTradePrepayOfflineReduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDTO {
    return r._offlineReducePrepayParam
}
