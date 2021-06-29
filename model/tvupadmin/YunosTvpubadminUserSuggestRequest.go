package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关联账户列表 APIRequest
yunos.tvpubadmin.user.suggest

获取关联账户列表
*/
type YunosTvpubadminUserSuggestRequest struct {
    model.Params

    // 关键词
    keyword   string 

}

func NewYunosTvpubadminUserSuggestRequest() *YunosTvpubadminUserSuggestRequest{
    return &YunosTvpubadminUserSuggestRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminUserSuggestRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.user.suggest"
}

func (r YunosTvpubadminUserSuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminUserSuggestRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r YunosTvpubadminUserSuggestRequest) GetKeyword() string {
    return r.keyword
}

