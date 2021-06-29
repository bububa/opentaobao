package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询全局弹窗 APIRequest
yunos.tvpubadmin.manage.dialog.findbyid

根据id查询全局弹窗
*/
type YunosTvpubadminManageDialogFindbyidRequest struct {
    model.Params

    // 全局弹窗id
    id   int64 

}

func NewYunosTvpubadminManageDialogFindbyidRequest() *YunosTvpubadminManageDialogFindbyidRequest{
    return &YunosTvpubadminManageDialogFindbyidRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminManageDialogFindbyidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.dialog.findbyid"
}

func (r YunosTvpubadminManageDialogFindbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminManageDialogFindbyidRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminManageDialogFindbyidRequest) GetId() int64 {
    return r.id
}

