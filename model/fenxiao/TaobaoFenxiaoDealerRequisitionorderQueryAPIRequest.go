package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest 按编号查询采购申请/经销采购单 API请求
// taobao.fenxiao.dealer.requisitionorder.query
//
// 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
type TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest struct {
	model.Params
	// 经销采购单编号。<br/>多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。
	_dealerOrderIds []int64
	// 多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
	_fields string
}

// NewTaobaoFenxiaoDealerRequisitionorderQueryRequest 初始化TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderQueryRequest() *TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest {
	return &TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDealerOrderIds is DealerOrderIds Setter
// 经销采购单编号。&lt;br/&gt;多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。
func (r *TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest) SetDealerOrderIds(_dealerOrderIds []int64) error {
	r._dealerOrderIds = _dealerOrderIds
	r.Set("dealer_order_ids", _dealerOrderIds)
	return nil
}

// GetDealerOrderIds DealerOrderIds Getter
func (r TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest) GetDealerOrderIds() []int64 {
	return r._dealerOrderIds
}

// SetFields is Fields Setter
// 多个字段用&#34;,&#34;分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
func (r *TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest) GetFields() string {
	return r._fields
}
