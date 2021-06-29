package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品下挂载的技能列表 APIRequest
taobao.ailab.aicloud.top.skils.list.new

星空平台提供的获取产品下挂载的技能列表新接口
*/
type TaobaoAilabAicloudTopSkilsListNewRequest struct {
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

func NewTaobaoAilabAicloudTopSkilsListNewRequest() *TaobaoAilabAicloudTopSkilsListNewRequest{
    return &TaobaoAilabAicloudTopSkilsListNewRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.skils.list.new"
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetQuery() string {
    return r.query
}

func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetType() string {
    return r.type
}

func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoAilabAicloudTopSkilsListNewRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoAilabAicloudTopSkilsListNewRequest) GetPageSize() int64 {
    return r.pageSize
}

