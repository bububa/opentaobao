package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道间库存规则设置接口 API请求
taobao.qimen.inventoryrule.create

渠道间库存规则设置
*/
type TaobaoQimenInventoryruleCreateRequest struct {
    model.Params
    // 
    _request   *RequestDO
}

// 初始化TaobaoQimenInventoryruleCreateRequest对象
func NewTaobaoQimenInventoryruleCreateRequest() *TaobaoQimenInventoryruleCreateRequest{
    return &TaobaoQimenInventoryruleCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryruleCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.inventoryrule.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryruleCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenInventoryruleCreateRequest) SetRequest(_request *RequestDO) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenInventoryruleCreateRequest) GetRequest() *RequestDO {
    return r._request
}
