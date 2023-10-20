package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundmessagegetAPIRequest openmall获取退款单留言 API请求
// taobao.openmall.refund.message.get
//
// openmall获取退款单留言
type TaobaoopenmallrefundmessagegetAPIRequest struct {
	model.Params
	// 分销者身份
	_distributor string
	// 翻页页码
	_pageNo int64
	// 翻页大小
	_pageSize int64
	// 退款单号
	_refundId int64
}

// NewTaobaoopenmallrefundmessagegetRequest 初始化TaobaoopenmallrefundmessagegetAPIRequest对象
func NewTaobaoopenmallrefundmessagegetRequest() *TaobaoopenmallrefundmessagegetAPIRequest {
	return &TaobaoopenmallrefundmessagegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallrefundmessagegetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.message.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallrefundmessagegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallrefundmessagegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者身份
func (r *TaobaoopenmallrefundmessagegetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmallrefundmessagegetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetPageNo is PageNo Setter
// 翻页页码
func (r *TaobaoopenmallrefundmessagegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoopenmallrefundmessagegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 翻页大小
func (r *TaobaoopenmallrefundmessagegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoopenmallrefundmessagegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TaobaoopenmallrefundmessagegetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoopenmallrefundmessagegetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
