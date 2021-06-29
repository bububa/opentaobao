package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单确认接口 API请求
taobao.qimen.returnorder.confirm

taobao.qimen.returnorder.confirm
*/
type TaobaoQimenReturnorderConfirmRequest struct {
    model.Params
    // 
    request   *ReturnOrderConfirmRequest
}

// 初始化TaobaoQimenReturnorderConfirmRequest对象
func NewTaobaoQimenReturnorderConfirmRequest() *TaobaoQimenReturnorderConfirmRequest{
    return &TaobaoQimenReturnorderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnorderConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.returnorder.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnorderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenReturnorderConfirmRequest) SetRequest(request *ReturnOrderConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenReturnorderConfirmRequest) GetRequest() *ReturnOrderConfirmRequest {
    return r.request
}
