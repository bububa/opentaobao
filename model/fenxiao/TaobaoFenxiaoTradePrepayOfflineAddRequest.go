package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下预存款流水增加 APIRequest
taobao.fenxiao.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加）
*/
type TaobaoFenxiaoTradePrepayOfflineAddRequest struct {
    model.Params

    // 增加流水
    offlineAddPrepayParam   *TopOfflineAddPrepayDto 

}

func NewTaobaoFenxiaoTradePrepayOfflineAddRequest() *TaobaoFenxiaoTradePrepayOfflineAddRequest{
    return &TaobaoFenxiaoTradePrepayOfflineAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoTradePrepayOfflineAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.trade.prepay.offline.add"
}

func (r TaobaoFenxiaoTradePrepayOfflineAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoTradePrepayOfflineAddRequest) SetOfflineAddPrepayParam(offlineAddPrepayParam *TopOfflineAddPrepayDto) error {
    r.offlineAddPrepayParam = offlineAddPrepayParam
    r.Set("offline_add_prepay_param", offlineAddPrepayParam)
    return nil
}

func (r TaobaoFenxiaoTradePrepayOfflineAddRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDto {
    return r.offlineAddPrepayParam
}

