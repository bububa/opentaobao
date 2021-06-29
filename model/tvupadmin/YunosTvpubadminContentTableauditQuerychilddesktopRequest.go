package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松查看小酷宝桌面坑位元数据列表 APIRequest
yunos.tvpubadmin.content.tableaudit.querychilddesktop

迎客松查看小酷宝桌面坑位元数据列表
*/
type YunosTvpubadminContentTableauditQuerychilddesktopRequest struct {
    model.Params

    // 小酷宝桌面坑位查询参数
    query   string 

}

func NewYunosTvpubadminContentTableauditQuerychilddesktopRequest() *YunosTvpubadminContentTableauditQuerychilddesktopRequest{
    return &YunosTvpubadminContentTableauditQuerychilddesktopRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentTableauditQuerychilddesktopRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.querychilddesktop"
}

func (r YunosTvpubadminContentTableauditQuerychilddesktopRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentTableauditQuerychilddesktopRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r YunosTvpubadminContentTableauditQuerychilddesktopRequest) GetQuery() string {
    return r.query
}

