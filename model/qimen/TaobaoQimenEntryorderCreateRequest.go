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
type TaobaoQimenEntryorderCreateRequest struct {
    model.Params
    // 
    request   *EntryOrderCreateRequest
}

// 初始化TaobaoQimenEntryorderCreateRequest对象
func NewTaobaoQimenEntryorderCreateRequest() *TaobaoQimenEntryorderCreateRequest{
    return &TaobaoQimenEntryorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.entryorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenEntryorderCreateRequest) SetRequest(request *EntryOrderCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenEntryorderCreateRequest) GetRequest() *EntryOrderCreateRequest {
    return r.request
}
