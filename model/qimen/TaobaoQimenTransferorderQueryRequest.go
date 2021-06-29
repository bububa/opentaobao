package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单查询 API请求
taobao.qimen.transferorder.query

调拨单查询
*/
type TaobaoQimenTransferorderQueryRequest struct {
    model.Params
    // 
    _request   *TaobaoQimenTransferorderQueryStruct
}

// 初始化TaobaoQimenTransferorderQueryRequest对象
func NewTaobaoQimenTransferorderQueryRequest() *TaobaoQimenTransferorderQueryRequest{
    return &TaobaoQimenTransferorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenTransferorderQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.transferorder.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenTransferorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenTransferorderQueryRequest) SetRequest(_request *TaobaoQimenTransferorderQueryStruct) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenTransferorderQueryRequest) GetRequest() *TaobaoQimenTransferorderQueryStruct {
    return r._request
}
