package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
货主仓库资源查询接口 API请求
taobao.qimen.warehouseinfo.query

货主仓库资源查询
*/
type TaobaoQimenWarehouseinfoQueryRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenWarehouseinfoQueryRequest对象
func NewTaobaoQimenWarehouseinfoQueryRequest() *TaobaoQimenWarehouseinfoQueryRequest{
    return &TaobaoQimenWarehouseinfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenWarehouseinfoQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.warehouseinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenWarehouseinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenWarehouseinfoQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenWarehouseinfoQueryRequest) GetRequest() *Request {
    return r.request
}
