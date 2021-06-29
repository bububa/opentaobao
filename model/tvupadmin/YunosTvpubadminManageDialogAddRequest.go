package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.add

新增全局弹窗
*/
type YunosTvpubadminManageDialogAddRequest struct {
    model.Params
    // 新增的全局弹窗json
    _dialogJson   string
}

// 初始化YunosTvpubadminManageDialogAddRequest对象
func NewYunosTvpubadminManageDialogAddRequest() *YunosTvpubadminManageDialogAddRequest{
    return &YunosTvpubadminManageDialogAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DialogJson Setter
// 新增的全局弹窗json
func (r *YunosTvpubadminManageDialogAddRequest) SetDialogJson(_dialogJson string) error {
    r._dialogJson = _dialogJson
    r.Set("dialog_json", _dialogJson)
    return nil
}

// DialogJson Getter
func (r YunosTvpubadminManageDialogAddRequest) GetDialogJson() string {
    return r._dialogJson
}