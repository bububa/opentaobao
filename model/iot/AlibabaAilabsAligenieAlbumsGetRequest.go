package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专辑详情 API请求
alibaba.ailabs.aligenie.albums.get

给予厂商查询专辑下的音频详情
*/
type AlibabaAilabsAligenieAlbumsGetAPIRequest struct {
    model.Params
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 专辑 id
    _param1   int64
    // 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
    _param2   int64
    // 每页数量（不超过50）
    _param3   int64
}

// 初始化AlibabaAilabsAligenieAlbumsGetAPIRequest对象
func NewAlibabaAilabsAligenieAlbumsGetRequest() *AlibabaAilabsAligenieAlbumsGetAPIRequest{
    return &AlibabaAilabsAligenieAlbumsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.albums.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *AlibabaAilabsAligenieAlbumsGetAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaAilabsAligenieAlbumsGetAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *AlibabaAilabsAligenieAlbumsGetAPIRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *AlibabaAilabsAligenieAlbumsGetAPIRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetExt() string {
    return r._ext
}
// Param1 Setter
// 专辑 id
func (r *AlibabaAilabsAligenieAlbumsGetAPIRequest) SetParam1(_param1 int64) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetParam1() int64 {
    return r._param1
}
// Param2 Setter
// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
func (r *AlibabaAilabsAligenieAlbumsGetAPIRequest) SetParam2(_param2 int64) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetParam2() int64 {
    return r._param2
}
// Param3 Setter
// 每页数量（不超过50）
func (r *AlibabaAilabsAligenieAlbumsGetAPIRequest) SetParam3(_param3 int64) error {
    r._param3 = _param3
    r.Set("param3", _param3)
    return nil
}

// Param3 Getter
func (r AlibabaAilabsAligenieAlbumsGetAPIRequest) GetParam3() int64 {
    return r._param3
}
