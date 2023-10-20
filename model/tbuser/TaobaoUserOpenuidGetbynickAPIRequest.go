package tbuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouseropenuidgetbynickAPIRequest 根据买家nick获取买家openuid API请求
// taobao.user.openuid.getbynick
//
// 根据买家nick获取买家openuid，最大查询30个
type TaobaouseropenuidgetbynickAPIRequest struct {
	model.Params
	// 买家nick列表
	_buyerNicks string
}

// NewTaobaouseropenuidgetbynickRequest 初始化TaobaouseropenuidgetbynickAPIRequest对象
func NewTaobaouseropenuidgetbynickRequest() *TaobaouseropenuidgetbynickAPIRequest {
	return &TaobaouseropenuidgetbynickAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouseropenuidgetbynickAPIRequest) GetApiMethodName() string {
	return "taobao.user.openuid.getbynick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouseropenuidgetbynickAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouseropenuidgetbynickAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNicks is BuyerNicks Setter
// 买家nick列表
func (r *TaobaouseropenuidgetbynickAPIRequest) SetBuyerNicks(_buyerNicks string) error {
	r._buyerNicks = _buyerNicks
	r.Set("buyer_nicks", _buyerNicks)
	return nil
}

// GetBuyerNicks BuyerNicks Getter
func (r TaobaouseropenuidgetbynickAPIRequest) GetBuyerNicks() string {
	return r._buyerNicks
}
