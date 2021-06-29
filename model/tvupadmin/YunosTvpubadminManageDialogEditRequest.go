package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑全局弹窗 APIRequest
yunos.tvpubadmin.manage.dialog.edit

编辑全局弹窗
*/
type YunosTvpubadminManageDialogEditRequest struct {
    model.Params

    // 待编辑的全局弹窗
    dialogJson   string 

}

func NewYunosTvpubadminManageDialogEditRequest() *YunosTvpubadminManageDialogEditRequest{
    return &YunosTvpubadminManageDialogEditRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageDialogEditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.edit"
}

func (r YunosTvpubadminManageDialogEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageDialogEditRequest) SetDialogJson(dialogJson string) error {
    r.dialogJson = dialogJson
    r.Set("dialog_json", dialogJson)
    return nil
}

func (r YunosTvpubadminManageDialogEditRequest) GetDialogJson() string {
    return r.dialogJson
}

