package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionlimitdiscountdetailgetAPIRequest 限时打折详情查询 API请求
// taobao.promotion.limitdiscount.detail.get
//
// 限时打折详情查询。查询出指定限时打折的对应商品记录信息。
type TaobaopromotionlimitdiscountdetailgetAPIRequest struct {
	model.Params
	// 限时打折ID。这个针对查询唯一限时打折情况。
	_limitDiscountId int64
}

// NewTaobaopromotionlimitdiscountdetailgetRequest 初始化TaobaopromotionlimitdiscountdetailgetAPIRequest对象
func NewTaobaopromotionlimitdiscountdetailgetRequest() *TaobaopromotionlimitdiscountdetailgetAPIRequest {
	return &TaobaopromotionlimitdiscountdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionlimitdiscountdetailgetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.limitdiscount.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionlimitdiscountdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionlimitdiscountdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLimitDiscountId is LimitDiscountId Setter
// 限时打折ID。这个针对查询唯一限时打折情况。
func (r *TaobaopromotionlimitdiscountdetailgetAPIRequest) SetLimitDiscountId(_limitDiscountId int64) error {
	r._limitDiscountId = _limitDiscountId
	r.Set("limit_discount_id", _limitDiscountId)
	return nil
}

// GetLimitDiscountId LimitDiscountId Getter
func (r TaobaopromotionlimitdiscountdetailgetAPIRequest) GetLimitDiscountId() int64 {
	return r._limitDiscountId
}
