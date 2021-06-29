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
type AlibabaAilabsAligenieAlbumsSearchRequest struct {
    model.Params
    // 账户体系隔离
    schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 扩展信息，用于存放APP类型等
    ext   string
    // 一级类目，如：儿童、新闻、商业财经有声书等
    param1   string
    // 二级类目，如儿童下有：儿歌、童谣、国学等
    param2   string
    // 搜索的单个专辑名称
    param3   string
    // 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
    param4   int64
    // 每页数量（不超过50）
    param5   int64
}

// 初始化AlibabaAilabsAligenieAlbumsSearchRequest对象
func NewAlibabaAilabsAligenieAlbumsSearchRequest() *AlibabaAilabsAligenieAlbumsSearchRequest{
    return &AlibabaAilabsAligenieAlbumsSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.albums.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetExt() string {
    return r.ext
}
// Param1 Setter
// 一级类目，如：儿童、新闻、商业财经有声书等
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// 二级类目，如儿童下有：儿歌、童谣、国学等
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetParam2() string {
    return r.param2
}
// Param3 Setter
// 搜索的单个专辑名称
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetParam3(param3 string) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

// Param3 Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetParam3() string {
    return r.param3
}
// Param4 Setter
// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetParam4(param4 int64) error {
    r.param4 = param4
    r.Set("param4", param4)
    return nil
}

// Param4 Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetParam4() int64 {
    return r.param4
}
// Param5 Setter
// 每页数量（不超过50）
func (r *AlibabaAilabsAligenieAlbumsSearchRequest) SetParam5(param5 int64) error {
    r.param5 = param5
    r.Set("param5", param5)
    return nil
}

// Param5 Getter
func (r AlibabaAilabsAligenieAlbumsSearchRequest) GetParam5() int64 {
    return r.param5
}
