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
type TaobaoChannelTradePrepayOfflineAddAPIRequest struct {
    model.Params
    // 增加流水
    _offlineAddPrepayParam   *TopOfflineAddPrepayDTO
}

// 初始化TaobaoChannelTradePrepayOfflineAddAPIRequest对象
func NewTaobaoChannelTradePrepayOfflineAddRequest() *TaobaoChannelTradePrepayOfflineAddAPIRequest{
    return &TaobaoChannelTradePrepayOfflineAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoChannelTradePrepayOfflineAddAPIRequest) GetApiMethodName() string {
    return "taobao.channel.trade.prepay.offline.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoChannelTradePrepayOfflineAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineAddPrepayParam Setter
// 增加流水
func (r *TaobaoChannelTradePrepayOfflineAddAPIRequest) SetOfflineAddPrepayParam(_offlineAddPrepayParam *TopOfflineAddPrepayDTO) error {
    r._offlineAddPrepayParam = _offlineAddPrepayParam
    r.Set("offline_add_prepay_param", _offlineAddPrepayParam)
    return nil
}

// OfflineAddPrepayParam Getter
func (r TaobaoChannelTradePrepayOfflineAddAPIRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDTO {
    return r._offlineAddPrepayParam
}
