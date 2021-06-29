package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取热词 API请求
taobao.ailab.aicloud.top.hotwords.get

获取ASR热词
*/
type TaobaoAilabAicloudTopHotwordsGetRequest struct {
    model.Params
    // 三方用户id
    userId   string
    // 业务类型
    bizClass   string
    // schemeKey
    schema   string
}

// 初始化TaobaoAilabAicloudTopHotwordsGetRequest对象
func NewTaobaoAilabAicloudTopHotwordsGetRequest() *TaobaoAilabAicloudTopHotwordsGetRequest{
    return &TaobaoAilabAicloudTopHotwordsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.hotwords.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 三方用户id
func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetUserId() string {
    return r.userId
}
// BizClass Setter
// 业务类型
func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetBizClass(bizClass string) error {
    r.bizClass = bizClass
    r.Set("biz_class", bizClass)
    return nil
}

// BizClass Getter
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetBizClass() string {
    return r.bizClass
}
// Schema Setter
// schemeKey
func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetSchema() string {
    return r.schema
}
