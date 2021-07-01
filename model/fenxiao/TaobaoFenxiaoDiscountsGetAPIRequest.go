package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDiscountsGetAPIRequest
获取折扣信息 API请求
taobao.fenxiao.discounts.get

查询折扣信息 */
type TaobaoFenxiaoDiscountsGetAPIRequest struct {
	model.Params
	// 折扣ID
	_discountId int64
	// 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
	_extFields string
}

// NewTaobaoFenxiaoDiscountsGetRequest 初始化TaobaoFenxiaoDiscountsGetAPIRequest对象
func NewTaobaoFenxiaoDiscountsGetRequest() *TaobaoFenxiaoDiscountsGetAPIRequest {
	return &TaobaoFenxiaoDiscountsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDiscountsGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.discounts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDiscountsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DiscountId Setter
// 折扣ID
func (r *TaobaoFenxiaoDiscountsGetAPIRequest) SetDiscountId(_discountId int64) error {
	r._discountId = _discountId
	r.Set("discount_id", _discountId)
	return nil
}

// Get DiscountId Getter
func (r TaobaoFenxiaoDiscountsGetAPIRequest) GetDiscountId() int64 {
	return r._discountId
}

// Set is ExtFields Setter
// 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
func (r *TaobaoFenxiaoDiscountsGetAPIRequest) SetExtFields(_extFields string) error {
	r._extFields = _extFields
	r.Set("ext_fields", _extFields)
	return nil
}

// Get ExtFields Getter
func (r TaobaoFenxiaoDiscountsGetAPIRequest) GetExtFields() string {
	return r._extFields
}
