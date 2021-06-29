package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单创建 API请求
taobao.qimen.transferorder.create

调拨单创建
*/
type TaobaoQimenTransferorderCreateRequest struct {
    model.Params
    // 
    request   *TaobaoQimenTransferorderCreateStruct
}

// 初始化TaobaoQimenTransferorderCreateRequest对象
func NewTaobaoQimenTransferorderCreateRequest() *TaobaoQimenTransferorderCreateRequest{
    return &TaobaoQimenTransferorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenTransferorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.transferorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenTransferorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenTransferorderCreateRequest) SetRequest(request *TaobaoQimenTransferorderCreateStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenTransferorderCreateRequest) GetRequest() *TaobaoQimenTransferorderCreateStruct {
    return r.request
}
