package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildRootnodeGetAPIRequest 获取少儿大厅根类目接口 API请求
// yunos.tvpubadmin.content.child.rootnode.get
//
// 通过此接口可获取少儿大厅根类目列表
type YunosTvpubadminContentChildRootnodeGetAPIRequest struct {
	model.Params
	// 是否需要首页目录
	_needHomePage bool
}

// NewYunosTvpubadminContentChildRootnodeGetRequest 初始化YunosTvpubadminContentChildRootnodeGetAPIRequest对象
func NewYunosTvpubadminContentChildRootnodeGetRequest() *YunosTvpubadminContentChildRootnodeGetAPIRequest {
	return &YunosTvpubadminContentChildRootnodeGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentChildRootnodeGetAPIRequest) Reset() {
	r._needHomePage = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildRootnodeGetAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.rootnode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildRootnodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentChildRootnodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNeedHomePage is NeedHomePage Setter
// 是否需要首页目录
func (r *YunosTvpubadminContentChildRootnodeGetAPIRequest) SetNeedHomePage(_needHomePage bool) error {
	r._needHomePage = _needHomePage
	r.Set("need_home_page", _needHomePage)
	return nil
}

// GetNeedHomePage NeedHomePage Getter
func (r YunosTvpubadminContentChildRootnodeGetAPIRequest) GetNeedHomePage() bool {
	return r._needHomePage
}

var poolYunosTvpubadminContentChildRootnodeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentChildRootnodeGetRequest()
	},
}

// GetYunosTvpubadminContentChildRootnodeGetRequest 从 sync.Pool 获取 YunosTvpubadminContentChildRootnodeGetAPIRequest
func GetYunosTvpubadminContentChildRootnodeGetAPIRequest() *YunosTvpubadminContentChildRootnodeGetAPIRequest {
	return poolYunosTvpubadminContentChildRootnodeGetAPIRequest.Get().(*YunosTvpubadminContentChildRootnodeGetAPIRequest)
}

// ReleaseYunosTvpubadminContentChildRootnodeGetAPIRequest 将 YunosTvpubadminContentChildRootnodeGetAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentChildRootnodeGetAPIRequest(v *YunosTvpubadminContentChildRootnodeGetAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentChildRootnodeGetAPIRequest.Put(v)
}
