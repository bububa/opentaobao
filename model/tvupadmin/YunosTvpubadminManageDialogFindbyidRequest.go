package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.findbyid

根据id查询全局弹窗
*/
type YunosTvpubadminManageDialogFindbyidAPIRequest struct {
    model.Params
    // 全局弹窗id
    _id   int64
}

// 初始化YunosTvpubadminManageDialogFindbyidAPIRequest对象
func NewYunosTvpubadminManageDialogFindbyidRequest() *YunosTvpubadminManageDialogFindbyidAPIRequest{
    return &YunosTvpubadminManageDialogFindbyidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogFindbyidAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.findbyid"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogFindbyidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 全局弹窗id
func (r *YunosTvpubadminManageDialogFindbyidAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminManageDialogFindbyidAPIRequest) GetId() int64 {
    return r._id
}
