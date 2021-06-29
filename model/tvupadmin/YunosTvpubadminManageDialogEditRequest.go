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
type YunosTvpubadminManageDialogEditRequest struct {
    model.Params
    // 待编辑的全局弹窗
    _dialogJson   string
}

// 初始化YunosTvpubadminManageDialogEditRequest对象
func NewYunosTvpubadminManageDialogEditRequest() *YunosTvpubadminManageDialogEditRequest{
    return &YunosTvpubadminManageDialogEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogEditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.edit"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DialogJson Setter
// 待编辑的全局弹窗
func (r *YunosTvpubadminManageDialogEditRequest) SetDialogJson(_dialogJson string) error {
    r._dialogJson = _dialogJson
    r.Set("dialog_json", _dialogJson)
    return nil
}

// DialogJson Getter
func (r YunosTvpubadminManageDialogEditRequest) GetDialogJson() string {
    return r._dialogJson
}
