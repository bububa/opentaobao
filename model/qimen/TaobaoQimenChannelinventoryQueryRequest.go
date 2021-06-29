package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 API请求
taobao.qimen.channelinventory.query

渠道库存查询
*/
type TaobaoQimenChannelinventoryQueryRequest struct {
    model.Params
    // 
    request   *RequestDO
}

// 初始化TaobaoQimenChannelinventoryQueryRequest对象
func NewTaobaoQimenChannelinventoryQueryRequest() *TaobaoQimenChannelinventoryQueryRequest{
    return &TaobaoQimenChannelinventoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenChannelinventoryQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.channelinventory.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenChannelinventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenChannelinventoryQueryRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenChannelinventoryQueryRequest) GetRequest() *RequestDO {
    return r.request
}
