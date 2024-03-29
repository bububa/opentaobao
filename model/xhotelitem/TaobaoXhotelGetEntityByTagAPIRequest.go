package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelGetEntityByTagAPIRequest 根据标签查询实体 API请求
// taobao.xhotel.get.entity.by.tag
//
// 根据标签查询实体
type TaobaoXhotelGetEntityByTagAPIRequest struct {
	model.Params
	// 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
	_tokenStr string
	// 标签
	_tag int64
}

// NewTaobaoXhotelGetEntityByTagRequest 初始化TaobaoXhotelGetEntityByTagAPIRequest对象
func NewTaobaoXhotelGetEntityByTagRequest() *TaobaoXhotelGetEntityByTagAPIRequest {
	return &TaobaoXhotelGetEntityByTagAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelGetEntityByTagAPIRequest) Reset() {
	r._tokenStr = ""
	r._tag = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.get.entity.by.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTokenStr is TokenStr Setter
// 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
func (r *TaobaoXhotelGetEntityByTagAPIRequest) SetTokenStr(_tokenStr string) error {
	r._tokenStr = _tokenStr
	r.Set("token_str", _tokenStr)
	return nil
}

// GetTokenStr TokenStr Getter
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetTokenStr() string {
	return r._tokenStr
}

// SetTag is Tag Setter
// 标签
func (r *TaobaoXhotelGetEntityByTagAPIRequest) SetTag(_tag int64) error {
	r._tag = _tag
	r.Set("tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetTag() int64 {
	return r._tag
}

var poolTaobaoXhotelGetEntityByTagAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelGetEntityByTagRequest()
	},
}

// GetTaobaoXhotelGetEntityByTagRequest 从 sync.Pool 获取 TaobaoXhotelGetEntityByTagAPIRequest
func GetTaobaoXhotelGetEntityByTagAPIRequest() *TaobaoXhotelGetEntityByTagAPIRequest {
	return poolTaobaoXhotelGetEntityByTagAPIRequest.Get().(*TaobaoXhotelGetEntityByTagAPIRequest)
}

// ReleaseTaobaoXhotelGetEntityByTagAPIRequest 将 TaobaoXhotelGetEntityByTagAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelGetEntityByTagAPIRequest(v *TaobaoXhotelGetEntityByTagAPIRequest) {
	v.Reset()
	poolTaobaoXhotelGetEntityByTagAPIRequest.Put(v)
}
