package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoreceiptstateupdateAPIRequest 服务工单状态更新 API请求
// tmall.aliauto.receipt.state.update
//
// 二轮车服务工单状态更新
type TmallaliautoreceiptstateupdateAPIRequest struct {
	model.Params
	// FINISH:服务完成
	_status string
	// 上门服务时指定门店ID
	_outerStoreId string
	// 服务工单id
	_receiptId int64
}

// NewTmallaliautoreceiptstateupdateRequest 初始化TmallaliautoreceiptstateupdateAPIRequest对象
func NewTmallaliautoreceiptstateupdateRequest() *TmallaliautoreceiptstateupdateAPIRequest {
	return &TmallaliautoreceiptstateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoreceiptstateupdateAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.receipt.state.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoreceiptstateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoreceiptstateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// FINISH:服务完成
func (r *TmallaliautoreceiptstateupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallaliautoreceiptstateupdateAPIRequest) GetStatus() string {
	return r._status
}

// SetOuterStoreId is OuterStoreId Setter
// 上门服务时指定门店ID
func (r *TmallaliautoreceiptstateupdateAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r TmallaliautoreceiptstateupdateAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetReceiptId is ReceiptId Setter
// 服务工单id
func (r *TmallaliautoreceiptstateupdateAPIRequest) SetReceiptId(_receiptId int64) error {
	r._receiptId = _receiptId
	r.Set("receipt_id", _receiptId)
	return nil
}

// GetReceiptId ReceiptId Getter
func (r TmallaliautoreceiptstateupdateAPIRequest) GetReceiptId() int64 {
	return r._receiptId
}
