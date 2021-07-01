package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户是否签署支付宝代扣协议 API请求
taobao.oc.ap.contractsigned.get

用户是否签署支付宝代扣协议
*/
type TaobaoOcApContractsignedGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoOcApContractsignedGetAPIRequest对象
func NewTaobaoOcApContractsignedGetRequest() *TaobaoOcApContractsignedGetAPIRequest{
    return &TaobaoOcApContractsignedGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcApContractsignedGetAPIRequest) GetApiMethodName() string {
    return "taobao.oc.ap.contractsigned.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcApContractsignedGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
