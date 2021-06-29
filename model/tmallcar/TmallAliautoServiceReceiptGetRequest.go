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
type TmallAliautoServiceReceiptGetRequest struct {
    model.Params
    // 工单号
    receiptId   int64
}

// 初始化TmallAliautoServiceReceiptGetRequest对象
func NewTmallAliautoServiceReceiptGetRequest() *TmallAliautoServiceReceiptGetRequest{
    return &TmallAliautoServiceReceiptGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoServiceReceiptGetRequest) GetApiMethodName() string {
    return "tmall.aliauto.service.receipt.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoServiceReceiptGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptId Setter
// 工单号
func (r *TmallAliautoServiceReceiptGetRequest) SetReceiptId(receiptId int64) error {
    r.receiptId = receiptId
    r.Set("receipt_id", receiptId)
    return nil
}

// ReceiptId Getter
func (r TmallAliautoServiceReceiptGetRequest) GetReceiptId() int64 {
    return r.receiptId
}
