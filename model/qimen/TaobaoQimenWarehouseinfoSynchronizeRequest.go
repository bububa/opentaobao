package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓库同步接口 API请求
taobao.qimen.warehouseinfo.synchronize

仓库同步接口
*/
type TaobaoQimenWarehouseinfoSynchronizeRequest struct {
    model.Params
    // 请求报文
    _request   *Request
}

// 初始化TaobaoQimenWarehouseinfoSynchronizeRequest对象
func NewTaobaoQimenWarehouseinfoSynchronizeRequest() *TaobaoQimenWarehouseinfoSynchronizeRequest{
    return &TaobaoQimenWarehouseinfoSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenWarehouseinfoSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.warehouseinfo.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenWarehouseinfoSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 请求报文
func (r *TaobaoQimenWarehouseinfoSynchronizeRequest) SetRequest(_request *Request) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenWarehouseinfoSynchronizeRequest) GetRequest() *Request {
    return r._request
}
