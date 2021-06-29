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
    ext   string
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 账户体系隔离字符串
    schema   string
    // 欲解绑的设备ID
    uuid   string
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
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetExt() string {
    return r.ext
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetUserId() string {
    return r.userId
}
// Schema Setter
// 账户体系隔离字符串
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetSchema() string {
    return r.schema
}
// Uuid Setter
// 欲解绑的设备ID
func (r *AlibabaAilabsAligenieDeviceUnbindRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAilabsAligenieDeviceUnbindRequest) GetUuid() string {
    return r.uuid
}
