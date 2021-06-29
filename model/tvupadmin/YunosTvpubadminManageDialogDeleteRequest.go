package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.delete

删除全局弹窗
*/
type YunosTvpubadminManageDialogDeleteRequest struct {
    model.Params
    // 全局弹窗id
    id   int64
}

// 初始化YunosTvpubadminManageDialogDeleteRequest对象
func NewYunosTvpubadminManageDialogDeleteRequest() *YunosTvpubadminManageDialogDeleteRequest{
    return &YunosTvpubadminManageDialogDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogDeleteRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.delete"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 全局弹窗id
func (r *YunosTvpubadminManageDialogDeleteRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r YunosTvpubadminManageDialogDeleteRequest) GetId() int64 {
    return r.id
}
