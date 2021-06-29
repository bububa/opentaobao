package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取弹窗列表 APIRequest
yunos.tvpubadmin.manage.dialog.list

分页获取弹窗配置列表
*/
type YunosTvpubadminManageDialogListRequest struct {
    model.Params

    // 查询的query参数
    query   string 

}

func NewYunosTvpubadminManageDialogListRequest() *YunosTvpubadminManageDialogListRequest{
    return &YunosTvpubadminManageDialogListRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageDialogListRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.list"
}

func (r YunosTvpubadminManageDialogListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageDialogListRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r YunosTvpubadminManageDialogListRequest) GetQuery() string {
    return r.query
}

