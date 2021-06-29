package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增全局弹窗 APIRequest
yunos.tvpubadmin.manage.dialog.add

新增全局弹窗
*/
type YunosTvpubadminManageDialogAddRequest struct {
    model.Params

    // 新增的全局弹窗json
    dialogJson   string 

}

func NewYunosTvpubadminManageDialogAddRequest() *YunosTvpubadminManageDialogAddRequest{
    return &YunosTvpubadminManageDialogAddRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageDialogAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.add"
}

func (r YunosTvpubadminManageDialogAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageDialogAddRequest) SetDialogJson(dialogJson string) error {
    r.dialogJson = dialogJson
    r.Set("dialog_json", dialogJson)
    return nil
}

func (r YunosTvpubadminManageDialogAddRequest) GetDialogJson() string {
    return r.dialogJson
}

