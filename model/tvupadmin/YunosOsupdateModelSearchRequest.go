package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机型检索 API请求
yunos.osupdate.model.search

机型检索
*/
type YunosOsupdateModelSearchRequest struct {
    model.Params
    // 应用ID
    appId   int64
    // 关键词
    name   string
}

// 初始化YunosOsupdateModelSearchRequest对象
func NewYunosOsupdateModelSearchRequest() *YunosOsupdateModelSearchRequest{
    return &YunosOsupdateModelSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateModelSearchRequest) GetApiMethodName() string {
    return "yunos.osupdate.model.search"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateModelSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 应用ID
func (r *YunosOsupdateModelSearchRequest) SetAppId(appId int64) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r YunosOsupdateModelSearchRequest) GetAppId() int64 {
    return r.appId
}
// Name Setter
// 关键词
func (r *YunosOsupdateModelSearchRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r YunosOsupdateModelSearchRequest) GetName() string {
    return r.name
}
