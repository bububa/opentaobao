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
type YunosOsupdateModelSearchAPIRequest struct {
    model.Params
    // 应用ID
    _appId   int64
    // 关键词
    _name   string
}

// 初始化YunosOsupdateModelSearchAPIRequest对象
func NewYunosOsupdateModelSearchRequest() *YunosOsupdateModelSearchAPIRequest{
    return &YunosOsupdateModelSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateModelSearchAPIRequest) GetApiMethodName() string {
    return "yunos.osupdate.model.search"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateModelSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 应用ID
func (r *YunosOsupdateModelSearchAPIRequest) SetAppId(_appId int64) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r YunosOsupdateModelSearchAPIRequest) GetAppId() int64 {
    return r._appId
}
// Name Setter
// 关键词
func (r *YunosOsupdateModelSearchAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r YunosOsupdateModelSearchAPIRequest) GetName() string {
    return r._name
}
