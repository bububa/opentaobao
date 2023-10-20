package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeDeleteAPIRequest 删除创意 API请求
// taobao.simba.creative.delete
//
// 删除一个创意
type TaobaoSimbaCreativeDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 创意Id
	_creativeId int64
}

// NewTaobaoSimbaCreativeDeleteRequest 初始化TaobaoSimbaCreativeDeleteAPIRequest对象
func NewTaobaoSimbaCreativeDeleteRequest() *TaobaoSimbaCreativeDeleteAPIRequest {
	return &TaobaoSimbaCreativeDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCreativeDeleteAPIRequest) Reset() {
	r._nick = ""
	r._creativeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creative.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetNick() string {
	return r._nick
}

// SetCreativeId is CreativeId Setter
// 创意Id
func (r *TaobaoSimbaCreativeDeleteAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaoSimbaCreativeDeleteAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}

var poolTaobaoSimbaCreativeDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCreativeDeleteRequest()
	},
}

// GetTaobaoSimbaCreativeDeleteRequest 从 sync.Pool 获取 TaobaoSimbaCreativeDeleteAPIRequest
func GetTaobaoSimbaCreativeDeleteAPIRequest() *TaobaoSimbaCreativeDeleteAPIRequest {
	return poolTaobaoSimbaCreativeDeleteAPIRequest.Get().(*TaobaoSimbaCreativeDeleteAPIRequest)
}

// ReleaseTaobaoSimbaCreativeDeleteAPIRequest 将 TaobaoSimbaCreativeDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCreativeDeleteAPIRequest(v *TaobaoSimbaCreativeDeleteAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCreativeDeleteAPIRequest.Put(v)
}
