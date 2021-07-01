package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv查询服务工单详情 API请求
tmall.aliauto.service.receipt.get

isv查询服务工单详情
*/
type TmallAliautoServiceReceiptGetAPIRequest struct {
    model.Params
    // 工单号
    _receiptId   int64
}

// 初始化TmallAliautoServiceReceiptGetAPIRequest对象
func NewTmallAliautoServiceReceiptGetRequest() *TmallAliautoServiceReceiptGetAPIRequest{
    return &TmallAliautoServiceReceiptGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoServiceReceiptGetAPIRequest) GetApiMethodName() string {
    return "tmall.aliauto.service.receipt.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoServiceReceiptGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptId Setter
// 工单号
func (r *TmallAliautoServiceReceiptGetAPIRequest) SetReceiptId(_receiptId int64) error {
    r._receiptId = _receiptId
    r.Set("receipt_id", _receiptId)
    return nil
}

// ReceiptId Getter
func (r TmallAliautoServiceReceiptGetAPIRequest) GetReceiptId() int64 {
    return r._receiptId
}
