package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户是否签署支付宝代扣协议 APIRequest
taobao.oc.ap.contractsigned.get

用户是否签署支付宝代扣协议
*/
type TaobaoOcApContractsignedGetRequest struct {
    model.Params

}

func NewTaobaoOcApContractsignedGetRequest() *TaobaoOcApContractsignedGetRequest{
    return &TaobaoOcApContractsignedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOcApContractsignedGetRequest) GetApiMethodName() string {
    return "taobao.oc.ap.contractsigned.get"
}

func (r TaobaoOcApContractsignedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


