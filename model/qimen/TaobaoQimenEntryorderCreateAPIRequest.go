package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单创建接口 API请求
taobao.qimen.entryorder.create

ERP调用接口，创建入库单;
*/
type TaobaoQimenEntryorderCreateAPIRequest struct {
    model.Params
    // 
    _request   *EntryOrderCreateRequest
}

// 初始化TaobaoQimenEntryorderCreateAPIRequest对象
func NewTaobaoQimenEntryorderCreateRequest() *TaobaoQimenEntryorderCreateAPIRequest{
    return &TaobaoQimenEntryorderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderCreateAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.entryorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenEntryorderCreateAPIRequest) SetRequest(_request *EntryOrderCreateRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenEntryorderCreateAPIRequest) GetRequest() *EntryOrderCreateRequest {
    return r._request
}
