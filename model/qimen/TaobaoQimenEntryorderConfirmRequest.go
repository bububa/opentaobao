package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单确认接口 API请求
taobao.qimen.entryorder.confirm

WMS调用接口，回传入库单信息;
*/
type TaobaoQimenEntryorderConfirmRequest struct {
    model.Params
    // 
    _request   *EntryOrderConfirmRequest
}

// 初始化TaobaoQimenEntryorderConfirmRequest对象
func NewTaobaoQimenEntryorderConfirmRequest() *TaobaoQimenEntryorderConfirmRequest{
    return &TaobaoQimenEntryorderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.entryorder.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenEntryorderConfirmRequest) SetRequest(_request *EntryOrderConfirmRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenEntryorderConfirmRequest) GetRequest() *EntryOrderConfirmRequest {
    return r._request
}
