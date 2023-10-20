package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbpromogetAPIRequest 民宿查询营销活动 API请求
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
type TaobaoxhotelbnbpromogetAPIRequest struct {
	model.Params
	// 外部活动code
	_outerActivityCode string
}

// NewTaobaoxhotelbnbpromogetRequest 初始化TaobaoxhotelbnbpromogetAPIRequest对象
func NewTaobaoxhotelbnbpromogetRequest() *TaobaoxhotelbnbpromogetAPIRequest {
	return &TaobaoxhotelbnbpromogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbpromogetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbpromogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbpromogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterActivityCode is OuterActivityCode Setter
// 外部活动code
func (r *TaobaoxhotelbnbpromogetAPIRequest) SetOuterActivityCode(_outerActivityCode string) error {
	r._outerActivityCode = _outerActivityCode
	r.Set("outer_activity_code", _outerActivityCode)
	return nil
}

// GetOuterActivityCode OuterActivityCode Getter
func (r TaobaoxhotelbnbpromogetAPIRequest) GetOuterActivityCode() string {
	return r._outerActivityCode
}
