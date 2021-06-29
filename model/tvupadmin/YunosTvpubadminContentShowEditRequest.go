package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
媒资节目信息修改 API请求
yunos.tvpubadmin.content.show.edit

供迎客松修改媒资节目信息
*/
type YunosTvpubadminContentShowEditRequest struct {
    model.Params
    // 请求入参，JSON格式
    data   string
}

// 初始化YunosTvpubadminContentShowEditRequest对象
func NewYunosTvpubadminContentShowEditRequest() *YunosTvpubadminContentShowEditRequest{
    return &YunosTvpubadminContentShowEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowEditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.edit"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Data Setter
// 请求入参，JSON格式
func (r *YunosTvpubadminContentShowEditRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r YunosTvpubadminContentShowEditRequest) GetData() string {
    return r.data
}
