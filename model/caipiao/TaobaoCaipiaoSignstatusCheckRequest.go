package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
检查用户是否签署支付宝代购协议 API请求
taobao.caipiao.signstatus.check

检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。
*/
type TaobaoCaipiaoSignstatusCheckRequest struct {
    model.Params
}

// 初始化TaobaoCaipiaoSignstatusCheckRequest对象
func NewTaobaoCaipiaoSignstatusCheckRequest() *TaobaoCaipiaoSignstatusCheckRequest{
    return &TaobaoCaipiaoSignstatusCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoSignstatusCheckRequest) GetApiMethodName() string {
    return "taobao.caipiao.signstatus.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoSignstatusCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
