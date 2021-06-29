package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品下挂载的技能列表 API请求
taobao.ailab.aicloud.top.skils.list.new

星空平台提供的获取产品下挂载的技能列表新接口
*/
type TaobaoAilabAicloudTopSkilsListNewRequest struct {
    model.Params
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // query(模糊匹配skillName)
    _query   string
    // type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
    _type   string
    // pageNo
    _pageNo   int64
    // pageSize
    _pageSize   int64
}

// 初始化TaobaoAilabAicloudTopSkilsListNewRequest对象
func NewTaobaoAilabAicloudTopSkilsListNewRequest() *TaobaoAilabAicloudTopSkilsListNewRequest{
    return &TaobaoAilabAicloudTopSkilsListNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.skils.list.new"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetExt() string {
    return r._ext
}
// Query Setter
// query(模糊匹配skillName)
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetQuery() string {
    return r._query
}
// Type Setter
// type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetType() string {
    return r._type
}
// PageNo Setter
// pageNo
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// pageSize
func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetPageSize() int64 {
    return r._pageSize
}