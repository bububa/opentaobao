package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductcatsgetAPIRequest 查询产品线列表 API请求
// taobao.fenxiao.productcats.get
//
// 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
type TaobaofenxiaoproductcatsgetAPIRequest struct {
	model.Params
	// 返回字段列表
	_fields string
}

// NewTaobaofenxiaoproductcatsgetRequest 初始化TaobaofenxiaoproductcatsgetAPIRequest对象
func NewTaobaofenxiaoproductcatsgetRequest() *TaobaofenxiaoproductcatsgetAPIRequest {
	return &TaobaofenxiaoproductcatsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductcatsgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcats.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductcatsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductcatsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段列表
func (r *TaobaofenxiaoproductcatsgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaofenxiaoproductcatsgetAPIRequest) GetFields() string {
	return r._fields
}
