package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务工单状态更新 API请求
tmall.aliauto.receipt.state.update

二轮车服务工单状态更新
*/
type TmallAliautoReceiptStateUpdateRequest struct {
    model.Params
    // FINISH:服务完成
    status   string
    // 服务工单id
    receiptId   int64
}

// 初始化TmallAliautoReceiptStateUpdateRequest对象
func NewTmallAliautoReceiptStateUpdateRequest() *TmallAliautoReceiptStateUpdateRequest{
    return &TmallAliautoReceiptStateUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoReceiptStateUpdateRequest) GetApiMethodName() string {
    return "tmall.aliauto.receipt.state.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoReceiptStateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// FINISH:服务完成
func (r *TmallAliautoReceiptStateUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TmallAliautoReceiptStateUpdateRequest) GetStatus() string {
    return r.status
}
// ReceiptId Setter
// 服务工单id
func (r *TmallAliautoReceiptStateUpdateRequest) SetReceiptId(receiptId int64) error {
    r.receiptId = receiptId
    r.Set("receipt_id", receiptId)
    return nil
}

// ReceiptId Getter
func (r TmallAliautoReceiptStateUpdateRequest) GetReceiptId() int64 {
    return r.receiptId
}
