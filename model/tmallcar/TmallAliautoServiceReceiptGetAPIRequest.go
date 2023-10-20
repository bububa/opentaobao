package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoservicereceiptgetAPIRequest isv查询服务工单详情 API请求
// tmall.aliauto.service.receipt.get
//
// isv查询服务工单详情
type TmallaliautoservicereceiptgetAPIRequest struct {
	model.Params
	// 工单号
	_receiptId int64
}

// NewTmallaliautoservicereceiptgetRequest 初始化TmallaliautoservicereceiptgetAPIRequest对象
func NewTmallaliautoservicereceiptgetRequest() *TmallaliautoservicereceiptgetAPIRequest {
	return &TmallaliautoservicereceiptgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoservicereceiptgetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.service.receipt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoservicereceiptgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoservicereceiptgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptId is ReceiptId Setter
// 工单号
func (r *TmallaliautoservicereceiptgetAPIRequest) SetReceiptId(_receiptId int64) error {
	r._receiptId = _receiptId
	r.Set("receipt_id", _receiptId)
	return nil
}

// GetReceiptId ReceiptId Getter
func (r TmallaliautoservicereceiptgetAPIRequest) GetReceiptId() int64 {
	return r._receiptId
}
