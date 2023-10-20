package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScRelationRefundAPIRequest 淘宝客-服务商-维权退款订单查询 API请求
// taobao.tbk.sc.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
type TaobaoTbkScRelationRefundAPIRequest struct {
	model.Params
	// 参数option
	_searchOption *TopApiRefundRptOption
}

// NewTaobaoTbkScRelationRefundRequest 初始化TaobaoTbkScRelationRefundAPIRequest对象
func NewTaobaoTbkScRelationRefundRequest() *TaobaoTbkScRelationRefundAPIRequest {
	return &TaobaoTbkScRelationRefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScRelationRefundAPIRequest) Reset() {
	r._searchOption = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScRelationRefundAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.relation.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScRelationRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScRelationRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchOption is SearchOption Setter
// 参数option
func (r *TaobaoTbkScRelationRefundAPIRequest) SetSearchOption(_searchOption *TopApiRefundRptOption) error {
	r._searchOption = _searchOption
	r.Set("search_option", _searchOption)
	return nil
}

// GetSearchOption SearchOption Getter
func (r TaobaoTbkScRelationRefundAPIRequest) GetSearchOption() *TopApiRefundRptOption {
	return r._searchOption
}

var poolTaobaoTbkScRelationRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScRelationRefundRequest()
	},
}

// GetTaobaoTbkScRelationRefundRequest 从 sync.Pool 获取 TaobaoTbkScRelationRefundAPIRequest
func GetTaobaoTbkScRelationRefundAPIRequest() *TaobaoTbkScRelationRefundAPIRequest {
	return poolTaobaoTbkScRelationRefundAPIRequest.Get().(*TaobaoTbkScRelationRefundAPIRequest)
}

// ReleaseTaobaoTbkScRelationRefundAPIRequest 将 TaobaoTbkScRelationRefundAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScRelationRefundAPIRequest(v *TaobaoTbkScRelationRefundAPIRequest) {
	v.Reset()
	poolTaobaoTbkScRelationRefundAPIRequest.Put(v)
}
