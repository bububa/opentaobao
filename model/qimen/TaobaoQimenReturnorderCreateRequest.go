package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单创建接口 API请求
taobao.qimen.returnorder.create

ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
*/
type TaobaoQimenReturnorderCreateRequest struct {
    model.Params
    // 
    _request   *ReturnOrderCreateRequest
}

// 初始化TaobaoQimenReturnorderCreateRequest对象
func NewTaobaoQimenReturnorderCreateRequest() *TaobaoQimenReturnorderCreateRequest{
    return &TaobaoQimenReturnorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.returnorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenReturnorderCreateRequest) SetRequest(_request *ReturnOrderCreateRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenReturnorderCreateRequest) GetRequest() *ReturnOrderCreateRequest {
    return r._request
}
