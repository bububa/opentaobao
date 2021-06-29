package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单查询接口 API请求
taobao.qimen.entryorder.query

ERP调用接口，查询入库单信息;
*/
type TaobaoQimenEntryorderQueryRequest struct {
    model.Params
    // 
    _request   *EntryOrderQueryRequest
}

// 初始化TaobaoQimenEntryorderQueryRequest对象
func NewTaobaoQimenEntryorderQueryRequest() *TaobaoQimenEntryorderQueryRequest{
    return &TaobaoQimenEntryorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.entryorder.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenEntryorderQueryRequest) SetRequest(_request *EntryOrderQueryRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenEntryorderQueryRequest) GetRequest() *EntryOrderQueryRequest {
    return r._request
}
