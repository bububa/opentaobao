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
type TaobaoFenxiaoTradePrepayOfflineReduceRequest struct {
    model.Params
    // 减少流水
    offlineReducePrepayParam   *TopOfflineReducePrepayDto
}

// 初始化TaobaoFenxiaoTradePrepayOfflineReduceRequest对象
func NewTaobaoFenxiaoTradePrepayOfflineReduceRequest() *TaobaoFenxiaoTradePrepayOfflineReduceRequest{
    return &TaobaoFenxiaoTradePrepayOfflineReduceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetApiMethodName() string {
    return "taobao.fenxiao.trade.prepay.offline.reduce"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineReducePrepayParam Setter
// 减少流水
func (r *TaobaoFenxiaoTradePrepayOfflineReduceRequest) SetOfflineReducePrepayParam(offlineReducePrepayParam *TopOfflineReducePrepayDto) error {
    r.offlineReducePrepayParam = offlineReducePrepayParam
    r.Set("offline_reduce_prepay_param", offlineReducePrepayParam)
    return nil
}

// OfflineReducePrepayParam Getter
func (r TaobaoFenxiaoTradePrepayOfflineReduceRequest) GetOfflineReducePrepayParam() *TopOfflineReducePrepayDto {
    return r.offlineReducePrepayParam
}
