package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询专辑 API请求
alibaba.ailabs.aligenie.albums.search

搜索类目下的专辑信息
*/
type AlibabaAilabsAligenieAlbumsSearchAPIRequest struct {
    model.Params
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 一级类目，如：儿童、新闻、商业财经有声书等
    _param1   string
    // 二级类目，如儿童下有：儿歌、童谣、国学等
    _param2   string
    // 搜索的单个专辑名称
    _param3   string
    // 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
    _param4   int64
    // 每页数量（不超过50）
    _param5   int64
}

// 初始化AlibabaAilabsAligenieAlbumsSearchAPIRequest对象
func NewAlibabaAilabsAligenieAlbumsSearchRequest() *AlibabaAilabsAligenieAlbumsSearchAPIRequest{
    return &AlibabaAilabsAligenieAlbumsSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.albums.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetExt() string {
    return r._ext
}
// Param1 Setter
// 一级类目，如：儿童、新闻、商业财经有声书等
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// 二级类目，如儿童下有：儿歌、童谣、国学等
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetParam2(_param2 string) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetParam2() string {
    return r._param2
}
// Param3 Setter
// 搜索的单个专辑名称
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetParam3(_param3 string) error {
    r._param3 = _param3
    r.Set("param3", _param3)
    return nil
}

// Param3 Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetParam3() string {
    return r._param3
}
// Param4 Setter
// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetParam4(_param4 int64) error {
    r._param4 = _param4
    r.Set("param4", _param4)
    return nil
}

// Param4 Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetParam4() int64 {
    return r._param4
}
// Param5 Setter
// 每页数量（不超过50）
func (r *AlibabaAilabsAligenieAlbumsSearchAPIRequest) SetParam5(_param5 int64) error {
    r._param5 = _param5
    r.Set("param5", _param5)
    return nil
}

// Param5 Getter
func (r AlibabaAilabsAligenieAlbumsSearchAPIRequest) GetParam5() int64 {
    return r._param5
}
