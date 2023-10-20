package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangereturngoodsagreeAPIRequest 卖家确认收货 API请求
// tmall.exchange.returngoods.agree
//
// 卖家确认收货
type TmallexchangereturngoodsagreeAPIRequest struct {
	model.Params
	// 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
	_fields []string
	// 换货单号ID
	_disputeId int64
}

// NewTmallexchangereturngoodsagreeRequest 初始化TmallexchangereturngoodsagreeAPIRequest对象
func NewTmallexchangereturngoodsagreeRequest() *TmallexchangereturngoodsagreeAPIRequest {
	return &TmallexchangereturngoodsagreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangereturngoodsagreeAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.returngoods.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangereturngoodsagreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangereturngoodsagreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
func (r *TmallexchangereturngoodsagreeAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangereturngoodsagreeAPIRequest) GetFields() []string {
	return r._fields
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangereturngoodsagreeAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangereturngoodsagreeAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}
