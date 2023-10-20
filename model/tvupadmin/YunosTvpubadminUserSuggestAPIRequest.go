package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminusersuggestAPIRequest 获取关联账户列表 API请求
// yunos.tvpubadmin.user.suggest
//
// 获取关联账户列表
type YunostvpubadminusersuggestAPIRequest struct {
	model.Params
	// 关键词
	_keyword string
}

// NewYunostvpubadminusersuggestRequest 初始化YunostvpubadminusersuggestAPIRequest对象
func NewYunostvpubadminusersuggestRequest() *YunostvpubadminusersuggestAPIRequest {
	return &YunostvpubadminusersuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminusersuggestAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.user.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminusersuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminusersuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键词
func (r *YunostvpubadminusersuggestAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r YunostvpubadminusersuggestAPIRequest) GetKeyword() string {
	return r._keyword
}
