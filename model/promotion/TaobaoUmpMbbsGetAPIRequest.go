package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpMbbsGetAPIRequest 获取营销积木块列表 API请求
// taobao.ump.mbbs.get
//
// 获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的
type TaobaoUmpMbbsGetAPIRequest struct {
	model.Params
	// 积木块类型。如果该字段为空表示查出所有类型的<br/>现在有且仅有如下几种：resource,condition,action,target
	_type string
}

// NewTaobaoUmpMbbsGetRequest 初始化TaobaoUmpMbbsGetAPIRequest对象
func NewTaobaoUmpMbbsGetRequest() *TaobaoUmpMbbsGetAPIRequest {
	return &TaobaoUmpMbbsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbsGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbbs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetType is Type Setter
// 积木块类型。如果该字段为空表示查出所有类型的<br/>现在有且仅有如下几种：resource,condition,action,target
func (r *TaobaoUmpMbbsGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoUmpMbbsGetAPIRequest) GetType() string {
	return r._type
}
