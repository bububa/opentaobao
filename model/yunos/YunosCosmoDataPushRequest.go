package yunos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
COSMO-PUSH模式数据接入 API请求
yunos.cosmo.data.push

YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
*/
type YunosCosmoDataPushRequest struct {
    model.Params
    // 业务方数据源唯一标识，由COSMO平台颁发
    _appId   string
    // 业务方推送数据，List结构的JSON序列化字符串
    _jsonModel   string
}

// 初始化YunosCosmoDataPushRequest对象
func NewYunosCosmoDataPushRequest() *YunosCosmoDataPushRequest{
    return &YunosCosmoDataPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosCosmoDataPushRequest) GetApiMethodName() string {
    return "yunos.cosmo.data.push"
}

// IRequest interface 方法, 获取API参数
func (r YunosCosmoDataPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 业务方数据源唯一标识，由COSMO平台颁发
func (r *YunosCosmoDataPushRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r YunosCosmoDataPushRequest) GetAppId() string {
    return r._appId
}
// JsonModel Setter
// 业务方推送数据，List结构的JSON序列化字符串
func (r *YunosCosmoDataPushRequest) SetJsonModel(_jsonModel string) error {
    r._jsonModel = _jsonModel
    r.Set("json_model", _jsonModel)
    return nil
}

// JsonModel Getter
func (r YunosCosmoDataPushRequest) GetJsonModel() string {
    return r._jsonModel
}
