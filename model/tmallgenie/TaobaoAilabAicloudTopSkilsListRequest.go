package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取硬件平台设备下挂载的技能列表 API请求
taobao.ailab.aicloud.top.skils.list

提供给在硬件平台接入设备的技能列表
*/
type TaobaoAilabAicloudTopSkilsListRequest struct {
    model.Params
    // 账户体系隔离
    schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 扩展信息，用于存放APP类型等
    ext   string
    // query(模糊匹配skillName)
    query   string
    // type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
    type   string
    // pageNo
    pageNo   int64
    // pageSize
    pageSize   int64
}

// 初始化TaobaoAilabAicloudTopSkilsListRequest对象
func NewTaobaoAilabAicloudTopSkilsListRequest() *TaobaoAilabAicloudTopSkilsListRequest{
    return &TaobaoAilabAicloudTopSkilsListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopSkilsListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.skils.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopSkilsListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetExt() string {
    return r.ext
}
// Query Setter
// query(模糊匹配skillName)
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetQuery() string {
    return r.query
}
// Type Setter
// type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetType() string {
    return r.type
}
// PageNo Setter
// pageNo
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// pageSize
func (r *TaobaoAilabAicloudTopSkilsListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAilabAicloudTopSkilsListRequest) GetPageSize() int64 {
    return r.pageSize
}
