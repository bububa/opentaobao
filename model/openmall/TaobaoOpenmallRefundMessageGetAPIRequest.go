package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundMessageGetAPIRequest openmall获取退款单留言 API请求
// taobao.openmall.refund.message.get
//
// openmall获取退款单留言
type TaobaoOpenmallRefundMessageGetAPIRequest struct {
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

// NewTaobaoOpenmallRefundMessageGetRequest 初始化TaobaoOpenmallRefundMessageGetAPIRequest对象
func NewTaobaoOpenmallRefundMessageGetRequest() *TaobaoOpenmallRefundMessageGetAPIRequest {
	return &TaobaoOpenmallRefundMessageGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundMessageGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.message.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundMessageGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDistributor is Distributor Setter
// 分销者身份
func (r *TaobaoOpenmallRefundMessageGetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallRefundMessageGetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetPageNo is PageNo Setter
// 翻页页码
func (r *TaobaoOpenmallRefundMessageGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoOpenmallRefundMessageGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 翻页大小
func (r *TaobaoOpenmallRefundMessageGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOpenmallRefundMessageGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TaobaoOpenmallRefundMessageGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoOpenmallRefundMessageGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
