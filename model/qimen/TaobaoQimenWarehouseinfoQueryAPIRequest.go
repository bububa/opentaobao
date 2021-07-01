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
type TaobaoQimenWarehouseinfoQueryAPIRequest struct {
    model.Params
    // 
    _request   *TaobaoQimenWarehouseinfoQueryRequest
}

// 初始化TaobaoQimenWarehouseinfoQueryAPIRequest对象
func NewTaobaoQimenWarehouseinfoQueryRequest() *TaobaoQimenWarehouseinfoQueryAPIRequest{
    return &TaobaoQimenWarehouseinfoQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenWarehouseinfoQueryAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.warehouseinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenWarehouseinfoQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenWarehouseinfoQueryAPIRequest) SetRequest(_request *TaobaoQimenWarehouseinfoQueryRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenWarehouseinfoQueryAPIRequest) GetRequest() *TaobaoQimenWarehouseinfoQueryRequest {
    return r._request
}
