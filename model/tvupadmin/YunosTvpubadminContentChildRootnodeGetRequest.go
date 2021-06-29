package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取少儿大厅根类目接口 APIRequest
yunos.tvpubadmin.content.child.rootnode.get

通过此接口可获取少儿大厅根类目列表
*/
type YunosTvpubadminContentChildRootnodeGetRequest struct {
    model.Params

    // 是否需要首页目录
    needHomePage   bool 

}

func NewYunosTvpubadminContentChildRootnodeGetRequest() *YunosTvpubadminContentChildRootnodeGetRequest{
    return &YunosTvpubadminContentChildRootnodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChildRootnodeGetRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.rootnode.get"
}

func (r YunosTvpubadminContentChildRootnodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentChildRootnodeGetRequest) SetNeedHomePage(needHomePage bool) error {
    r.needHomePage = needHomePage
    r.Set("need_home_page", needHomePage)
    return nil
}

func (r YunosTvpubadminContentChildRootnodeGetRequest) GetNeedHomePage() bool {
    return r.needHomePage
}

