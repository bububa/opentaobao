package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关联账户列表 API请求
yunos.tvpubadmin.user.suggest

获取关联账户列表
*/
type YunosTvpubadminUserSuggestRequest struct {
    model.Params
    // 关键词
    _keyword   string
}

// 初始化YunosTvpubadminUserSuggestRequest对象
func NewYunosTvpubadminUserSuggestRequest() *YunosTvpubadminUserSuggestRequest{
    return &YunosTvpubadminUserSuggestRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserSuggestRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.user.suggest"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserSuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 关键词
func (r *YunosTvpubadminUserSuggestRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r YunosTvpubadminUserSuggestRequest) GetKeyword() string {
    return r._keyword
}
