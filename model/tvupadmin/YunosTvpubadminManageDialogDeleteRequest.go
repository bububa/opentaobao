package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除全局弹窗 APIRequest
yunos.tvpubadmin.manage.dialog.delete

删除全局弹窗
*/
type YunosTvpubadminManageDialogDeleteRequest struct {
    model.Params

    // 全局弹窗id
    id   int64 

}

func NewYunosTvpubadminManageDialogDeleteRequest() *YunosTvpubadminManageDialogDeleteRequest{
    return &YunosTvpubadminManageDialogDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageDialogDeleteRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.delete"
}

func (r YunosTvpubadminManageDialogDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageDialogDeleteRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminManageDialogDeleteRequest) GetId() int64 {
    return r.id
}

