package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（增加） API请求
taobao.channel.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加）
*/
type TaobaoChannelTradePrepayOfflineAddRequest struct {
    model.Params
    // 增加流水
    _offlineAddPrepayParam   *TopOfflineAddPrepayDto
}

// 初始化TaobaoChannelTradePrepayOfflineAddRequest对象
func NewTaobaoChannelTradePrepayOfflineAddRequest() *TaobaoChannelTradePrepayOfflineAddRequest{
    return &TaobaoChannelTradePrepayOfflineAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoChannelTradePrepayOfflineAddRequest) GetApiMethodName() string {
    return "taobao.channel.trade.prepay.offline.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoChannelTradePrepayOfflineAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineAddPrepayParam Setter
// 增加流水
func (r *TaobaoChannelTradePrepayOfflineAddRequest) SetOfflineAddPrepayParam(_offlineAddPrepayParam *TopOfflineAddPrepayDto) error {
    r._offlineAddPrepayParam = _offlineAddPrepayParam
    r.Set("offline_add_prepay_param", _offlineAddPrepayParam)
    return nil
}

// OfflineAddPrepayParam Getter
func (r TaobaoChannelTradePrepayOfflineAddRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDto {
    return r._offlineAddPrepayParam
}
