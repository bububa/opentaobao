package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.edit

编辑全局弹窗
*/
type YunosTvpubadminManageDialogEditAPIRequest struct {
    model.Params
    // 待编辑的全局弹窗
    _dialogJson   string
}

// 初始化YunosTvpubadminManageDialogEditAPIRequest对象
func NewYunosTvpubadminManageDialogEditRequest() *YunosTvpubadminManageDialogEditAPIRequest{
    return &YunosTvpubadminManageDialogEditAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogEditAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.edit"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogEditAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DialogJson Setter
// 待编辑的全局弹窗
func (r *YunosTvpubadminManageDialogEditAPIRequest) SetDialogJson(_dialogJson string) error {
    r._dialogJson = _dialogJson
    r.Set("dialog_json", _dialogJson)
    return nil
}

// DialogJson Getter
func (r YunosTvpubadminManageDialogEditAPIRequest) GetDialogJson() string {
    return r._dialogJson
}
