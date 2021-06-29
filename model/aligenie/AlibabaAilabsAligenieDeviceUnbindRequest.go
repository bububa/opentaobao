package aligenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备解绑操作接口 API请求
alibaba.ailabs.aligenie.device.unbind

让开发者能根据设备ID进行解绑操作的接口
*/
type AlibabaAilabsAligenieDeviceUnbindRequest struct {
    model.Params
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 账户体系隔离字符串
    _schema   string
    // 欲解绑的设备ID
    _uuid   string
}

// 初始化AlibabaAilabsAligenieDeviceUnbindRequest对象
func NewAlibabaAilabsAligenieDeviceUnbindRequest() *AlibabaAilabsAligenieDeviceUnbindRequest{
    return &AlibabaAilabsAligenieDeviceUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.device.unbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetExt() string {
    return r._ext
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetUserId() string {
    return r._userId
}
// Schema Setter
// 账户体系隔离字符串
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetSchema() string {
    return r._schema
}
// Uuid Setter
// 欲解绑的设备ID
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetUuid() string {
    return r._uuid
}
