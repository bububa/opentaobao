package fans

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFansArenaPushAPIRequest 消息推送 API请求
// tmall.fans.arena.push
//
// 超级擂台消息推送
type TmallFansArenaPushAPIRequest struct {
	model.Params
	// 推送列表
	_pushList []PushMessageParamDo
}

// NewTmallFansArenaPushRequest 初始化TmallFansArenaPushAPIRequest对象
func NewTmallFansArenaPushRequest() *TmallFansArenaPushAPIRequest {
	return &TmallFansArenaPushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFansArenaPushAPIRequest) Reset() {
	r._pushList = r._pushList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFansArenaPushAPIRequest) GetApiMethodName() string {
	return "tmall.fans.arena.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFansArenaPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFansArenaPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushList is PushList Setter
// 推送列表
func (r *TmallFansArenaPushAPIRequest) SetPushList(_pushList []PushMessageParamDo) error {
	r._pushList = _pushList
	r.Set("push_list", _pushList)
	return nil
}

// GetPushList PushList Getter
func (r TmallFansArenaPushAPIRequest) GetPushList() []PushMessageParamDo {
	return r._pushList
}

var poolTmallFansArenaPushAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFansArenaPushRequest()
	},
}

// GetTmallFansArenaPushRequest 从 sync.Pool 获取 TmallFansArenaPushAPIRequest
func GetTmallFansArenaPushAPIRequest() *TmallFansArenaPushAPIRequest {
	return poolTmallFansArenaPushAPIRequest.Get().(*TmallFansArenaPushAPIRequest)
}

// ReleaseTmallFansArenaPushAPIRequest 将 TmallFansArenaPushAPIRequest 放入 sync.Pool
func ReleaseTmallFansArenaPushAPIRequest(v *TmallFansArenaPushAPIRequest) {
	v.Reset()
	poolTmallFansArenaPushAPIRequest.Put(v)
}
