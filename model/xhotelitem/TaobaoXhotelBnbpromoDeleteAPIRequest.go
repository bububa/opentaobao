package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbpromodeleteAPIRequest 民宿卖家活动删除 API请求
// taobao.xhotel.bnbpromo.delete
//
// 民宿删除营销活动
type TaobaoxhotelbnbpromodeleteAPIRequest struct {
	model.Params
	// 外部活动code
	_outerActivityCode string
}

// NewTaobaoxhotelbnbpromodeleteRequest 初始化TaobaoxhotelbnbpromodeleteAPIRequest对象
func NewTaobaoxhotelbnbpromodeleteRequest() *TaobaoxhotelbnbpromodeleteAPIRequest {
	return &TaobaoxhotelbnbpromodeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbpromodeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbpromodeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbpromodeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterActivityCode is OuterActivityCode Setter
// 外部活动code
func (r *TaobaoxhotelbnbpromodeleteAPIRequest) SetOuterActivityCode(_outerActivityCode string) error {
	r._outerActivityCode = _outerActivityCode
	r.Set("outer_activity_code", _outerActivityCode)
	return nil
}

// GetOuterActivityCode OuterActivityCode Getter
func (r TaobaoxhotelbnbpromodeleteAPIRequest) GetOuterActivityCode() string {
	return r._outerActivityCode
}
