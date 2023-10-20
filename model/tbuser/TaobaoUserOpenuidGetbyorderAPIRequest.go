package tbuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouseropenuidgetbyorderAPIRequest 根据订单获取买家openuid API请求
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
type TaobaouseropenuidgetbyorderAPIRequest struct {
	model.Params
	// 买家订单列表
	_tidInfos *BuyerOrder
}

// NewTaobaouseropenuidgetbyorderRequest 初始化TaobaouseropenuidgetbyorderAPIRequest对象
func NewTaobaouseropenuidgetbyorderRequest() *TaobaouseropenuidgetbyorderAPIRequest {
	return &TaobaouseropenuidgetbyorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouseropenuidgetbyorderAPIRequest) GetApiMethodName() string {
	return "taobao.user.openuid.getbyorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouseropenuidgetbyorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouseropenuidgetbyorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTidInfos is TidInfos Setter
// 买家订单列表
func (r *TaobaouseropenuidgetbyorderAPIRequest) SetTidInfos(_tidInfos *BuyerOrder) error {
	r._tidInfos = _tidInfos
	r.Set("tid_infos", _tidInfos)
	return nil
}

// GetTidInfos TidInfos Getter
func (r TaobaouseropenuidgetbyorderAPIRequest) GetTidInfos() *BuyerOrder {
	return r._tidInfos
}
