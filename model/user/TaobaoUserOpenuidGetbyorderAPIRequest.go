package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenuidGetbyorderAPIRequest 根据订单获取买家openuid API请求
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
type TaobaoUserOpenuidGetbyorderAPIRequest struct {
	model.Params
	// 买家订单列表
	_tidInfos *BuyerOrder
}

// NewTaobaoUserOpenuidGetbyorderRequest 初始化TaobaoUserOpenuidGetbyorderAPIRequest对象
func NewTaobaoUserOpenuidGetbyorderRequest() *TaobaoUserOpenuidGetbyorderAPIRequest {
	return &TaobaoUserOpenuidGetbyorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserOpenuidGetbyorderAPIRequest) GetApiMethodName() string {
	return "taobao.user.openuid.getbyorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserOpenuidGetbyorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTidInfos is TidInfos Setter
// 买家订单列表
func (r *TaobaoUserOpenuidGetbyorderAPIRequest) SetTidInfos(_tidInfos *BuyerOrder) error {
	r._tidInfos = _tidInfos
	r.Set("tid_infos", _tidInfos)
	return nil
}

// GetTidInfos TidInfos Getter
func (r TaobaoUserOpenuidGetbyorderAPIRequest) GetTidInfos() *BuyerOrder {
	return r._tidInfos
}
