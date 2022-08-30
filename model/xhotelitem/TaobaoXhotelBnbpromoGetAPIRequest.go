package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoGetAPIRequest 民宿查询营销活动 API请求
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
type TaobaoXhotelBnbpromoGetAPIRequest struct {
	model.Params
	// 外部活动code
	_outerActivityCode string
}

// NewTaobaoXhotelBnbpromoGetRequest 初始化TaobaoXhotelBnbpromoGetAPIRequest对象
func NewTaobaoXhotelBnbpromoGetRequest() *TaobaoXhotelBnbpromoGetAPIRequest {
	return &TaobaoXhotelBnbpromoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterActivityCode is OuterActivityCode Setter
// 外部活动code
func (r *TaobaoXhotelBnbpromoGetAPIRequest) SetOuterActivityCode(_outerActivityCode string) error {
	r._outerActivityCode = _outerActivityCode
	r.Set("outer_activity_code", _outerActivityCode)
	return nil
}

// GetOuterActivityCode OuterActivityCode Getter
func (r TaobaoXhotelBnbpromoGetAPIRequest) GetOuterActivityCode() string {
	return r._outerActivityCode
}
