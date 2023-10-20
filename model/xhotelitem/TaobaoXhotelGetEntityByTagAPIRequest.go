package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelgetentitybytagAPIRequest 根据标签查询实体 API请求
// taobao.xhotel.get.entity.by.tag
//
// 根据标签查询实体
type TaobaoxhotelgetentitybytagAPIRequest struct {
	model.Params
	// 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
	_tokenStr string
	// 标签
	_tag int64
}

// NewTaobaoxhotelgetentitybytagRequest 初始化TaobaoxhotelgetentitybytagAPIRequest对象
func NewTaobaoxhotelgetentitybytagRequest() *TaobaoxhotelgetentitybytagAPIRequest {
	return &TaobaoxhotelgetentitybytagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelgetentitybytagAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.get.entity.by.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelgetentitybytagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelgetentitybytagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTokenStr is TokenStr Setter
// 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
func (r *TaobaoxhotelgetentitybytagAPIRequest) SetTokenStr(_tokenStr string) error {
	r._tokenStr = _tokenStr
	r.Set("token_str", _tokenStr)
	return nil
}

// GetTokenStr TokenStr Getter
func (r TaobaoxhotelgetentitybytagAPIRequest) GetTokenStr() string {
	return r._tokenStr
}

// SetTag is Tag Setter
// 标签
func (r *TaobaoxhotelgetentitybytagAPIRequest) SetTag(_tag int64) error {
	r._tag = _tag
	r.Set("tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r TaobaoxhotelgetentitybytagAPIRequest) GetTag() int64 {
	return r._tag
}
