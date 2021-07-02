package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserSuggestAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.user.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserSuggestAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Keyword Setter
// 关键词
func (r *YunosTvpubadminUserSuggestAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// Get Keyword Getter
func (r YunosTvpubadminUserSuggestAPIRequest) GetKeyword() string {
	return r._keyword
}
