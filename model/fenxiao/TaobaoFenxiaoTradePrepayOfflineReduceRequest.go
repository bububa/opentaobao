package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（减少） API请求
taobao.fenxiao.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少）
*/
type TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest struct {
    model.Params
    // 减少流水
    _offlineReducePrepayParam   *TopOfflineReducePrepayDTO
}

// 初始化TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest对象
func NewTaobaoFenxiaoTradePrepayOfflineReduceRequest() *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest{
    return &TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.trade.prepay.offline.reduce"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) SetOfflineReducePrepayParam(_offlineReducePrepayParam *TopOfflineReducePrepayDTO) error {
    r._offlineReducePrepayParam = _offlineReducePrepayParam
    r.Set("offline_reduce_prepay_param", _offlineReducePrepayParam)
    return nil
}

// OfflineReducePrepayParam Getter
func (r TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDTO {
    return r._offlineReducePrepayParam
}
