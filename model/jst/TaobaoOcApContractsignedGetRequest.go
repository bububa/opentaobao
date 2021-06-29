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
type TaobaoOcApContractsignedGetRequest struct {
    model.Params
}

// 初始化TaobaoOcApContractsignedGetRequest对象
func NewTaobaoOcApContractsignedGetRequest() *TaobaoOcApContractsignedGetRequest{
    return &TaobaoOcApContractsignedGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcApContractsignedGetRequest) GetApiMethodName() string {
    return "taobao.oc.ap.contractsigned.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcApContractsignedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
