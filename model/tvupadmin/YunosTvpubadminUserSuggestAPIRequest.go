package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminUserSuggestAPIRequest 获取关联账户列表 API请求
// yunos.tvpubadmin.user.suggest
//
// 获取关联账户列表
type YunosTvpubadminUserSuggestAPIRequest struct {
	model.Params
	// 关键词
	_keyword string
}

// NewYunosTvpubadminUserSuggestRequest 初始化YunosTvpubadminUserSuggestAPIRequest对象
func NewYunosTvpubadminUserSuggestRequest() *YunosTvpubadminUserSuggestAPIRequest {
	return &YunosTvpubadminUserSuggestAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminUserSuggestAPIRequest) Reset() {
	r._keyword = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserSuggestAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.user.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserSuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminUserSuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键词
func (r *YunosTvpubadminUserSuggestAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r YunosTvpubadminUserSuggestAPIRequest) GetKeyword() string {
	return r._keyword
}

var poolYunosTvpubadminUserSuggestAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminUserSuggestRequest()
	},
}

// GetYunosTvpubadminUserSuggestRequest 从 sync.Pool 获取 YunosTvpubadminUserSuggestAPIRequest
func GetYunosTvpubadminUserSuggestAPIRequest() *YunosTvpubadminUserSuggestAPIRequest {
	return poolYunosTvpubadminUserSuggestAPIRequest.Get().(*YunosTvpubadminUserSuggestAPIRequest)
}

// ReleaseYunosTvpubadminUserSuggestAPIRequest 将 YunosTvpubadminUserSuggestAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminUserSuggestAPIRequest(v *YunosTvpubadminUserSuggestAPIRequest) {
	v.Reset()
	poolYunosTvpubadminUserSuggestAPIRequest.Put(v)
}
