package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodiscountsgetAPIRequest 获取折扣信息 API请求
// taobao.fenxiao.discounts.get
//
// 查询折扣信息
type TaobaofenxiaodiscountsgetAPIRequest struct {
	model.Params
	// 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
	_extFields string
	// 折扣ID
	_discountId int64
}

// NewTaobaofenxiaodiscountsgetRequest 初始化TaobaofenxiaodiscountsgetAPIRequest对象
func NewTaobaofenxiaodiscountsgetRequest() *TaobaofenxiaodiscountsgetAPIRequest {
	return &TaobaofenxiaodiscountsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodiscountsgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.discounts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodiscountsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodiscountsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtFields is ExtFields Setter
// 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
func (r *TaobaofenxiaodiscountsgetAPIRequest) SetExtFields(_extFields string) error {
	r._extFields = _extFields
	r.Set("ext_fields", _extFields)
	return nil
}

// GetExtFields ExtFields Getter
func (r TaobaofenxiaodiscountsgetAPIRequest) GetExtFields() string {
	return r._extFields
}

// SetDiscountId is DiscountId Setter
// 折扣ID
func (r *TaobaofenxiaodiscountsgetAPIRequest) SetDiscountId(_discountId int64) error {
	r._discountId = _discountId
	r.Set("discount_id", _discountId)
	return nil
}

// GetDiscountId DiscountId Getter
func (r TaobaofenxiaodiscountsgetAPIRequest) GetDiscountId() int64 {
	return r._discountId
}
