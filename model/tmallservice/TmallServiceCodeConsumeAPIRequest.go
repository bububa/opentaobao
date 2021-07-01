package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceCodeConsumeAPIRequest
天猫服务平台服务核销 API请求
tmall.service.code.consume

天猫服务平台－服务核销 */
type TmallServiceCodeConsumeAPIRequest struct {
	model.Params
	// 核销码
	_consumeCode string
	// 核销帐号
	_operatorNick string
	// 门店id
	_shopId string
}

// NewTmallServiceCodeConsumeRequest 初始化TmallServiceCodeConsumeAPIRequest对象
func NewTmallServiceCodeConsumeRequest() *TmallServiceCodeConsumeAPIRequest {
	return &TmallServiceCodeConsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceCodeConsumeAPIRequest) GetApiMethodName() string {
	return "tmall.service.code.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceCodeConsumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ConsumeCode Setter
// 核销码
func (r *TmallServiceCodeConsumeAPIRequest) SetConsumeCode(_consumeCode string) error {
	r._consumeCode = _consumeCode
	r.Set("consume_code", _consumeCode)
	return nil
}

// Get ConsumeCode Getter
func (r TmallServiceCodeConsumeAPIRequest) GetConsumeCode() string {
	return r._consumeCode
}

// Set is OperatorNick Setter
// 核销帐号
func (r *TmallServiceCodeConsumeAPIRequest) SetOperatorNick(_operatorNick string) error {
	r._operatorNick = _operatorNick
	r.Set("operator_nick", _operatorNick)
	return nil
}

// Get OperatorNick Getter
func (r TmallServiceCodeConsumeAPIRequest) GetOperatorNick() string {
	return r._operatorNick
}

// Set is ShopId Setter
// 门店id
func (r *TmallServiceCodeConsumeAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r TmallServiceCodeConsumeAPIRequest) GetShopId() string {
	return r._shopId
}
