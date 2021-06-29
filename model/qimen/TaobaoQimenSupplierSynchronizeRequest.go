package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商同步接口 API请求
taobao.qimen.supplier.synchronize

这个接口用来同步供应商信息
*/
type TaobaoQimenSupplierSynchronizeRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenSupplierSynchronizeRequest对象
func NewTaobaoQimenSupplierSynchronizeRequest() *TaobaoQimenSupplierSynchronizeRequest{
    return &TaobaoQimenSupplierSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenSupplierSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.supplier.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenSupplierSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenSupplierSynchronizeRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenSupplierSynchronizeRequest) GetRequest() *Request {
    return r.request
}
