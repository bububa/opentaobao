package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoDeleteAPIRequest 民宿卖家活动删除 API请求
// taobao.xhotel.bnbpromo.delete
//
// 民宿删除营销活动
type TaobaoXhotelBnbpromoDeleteAPIRequest struct {
	model.Params
	// 外部活动code
	_outerActivityCode string
}

// NewTaobaoXhotelBnbpromoDeleteRequest 初始化TaobaoXhotelBnbpromoDeleteAPIRequest对象
func NewTaobaoXhotelBnbpromoDeleteRequest() *TaobaoXhotelBnbpromoDeleteAPIRequest {
	return &TaobaoXhotelBnbpromoDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterActivityCode is OuterActivityCode Setter
// 外部活动code
func (r *TaobaoXhotelBnbpromoDeleteAPIRequest) SetOuterActivityCode(_outerActivityCode string) error {
	r._outerActivityCode = _outerActivityCode
	r.Set("outer_activity_code", _outerActivityCode)
	return nil
}

// GetOuterActivityCode OuterActivityCode Getter
func (r TaobaoXhotelBnbpromoDeleteAPIRequest) GetOuterActivityCode() string {
	return r._outerActivityCode
}
