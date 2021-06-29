package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下预存款流水增加 API请求
taobao.fenxiao.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加）
*/
type TaobaoFenxiaoTradePrepayOfflineAddRequest struct {
    model.Params
    // 增加流水
    _offlineAddPrepayParam   *TopOfflineAddPrepayDTO
}

// 初始化TaobaoFenxiaoTradePrepayOfflineAddRequest对象
func NewTaobaoFenxiaoTradePrepayOfflineAddRequest() *TaobaoFenxiaoTradePrepayOfflineAddRequest{
    return &TaobaoFenxiaoTradePrepayOfflineAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoTradePrepayOfflineAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.trade.prepay.offline.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoTradePrepayOfflineAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineAddPrepayParam Setter
// 增加流水
func (r *TaobaoFenxiaoTradePrepayOfflineAddRequest) SetOfflineAddPrepayParam(_offlineAddPrepayParam *TopOfflineAddPrepayDTO) error {
    r._offlineAddPrepayParam = _offlineAddPrepayParam
    r.Set("offline_add_prepay_param", _offlineAddPrepayParam)
    return nil
}

// OfflineAddPrepayParam Getter
func (r TaobaoFenxiaoTradePrepayOfflineAddRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDTO {
    return r._offlineAddPrepayParam
}
