package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemoutercodegetAPIRequest 根据outerCode查询商品 API请求
// taobao.scitem.outercode.get
//
// 根据outerCode查询商品
type TaobaoscitemoutercodegetAPIRequest struct {
	model.Params
	// 商品编码
	_outerCode string
}

// NewTaobaoscitemoutercodegetRequest 初始化TaobaoscitemoutercodegetAPIRequest对象
func NewTaobaoscitemoutercodegetRequest() *TaobaoscitemoutercodegetAPIRequest {
	return &TaobaoscitemoutercodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoscitemoutercodegetAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.outercode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoscitemoutercodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoscitemoutercodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// 商品编码
func (r *TaobaoscitemoutercodegetAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TaobaoscitemoutercodegetAPIRequest) GetOuterCode() string {
	return r._outerCode
}
