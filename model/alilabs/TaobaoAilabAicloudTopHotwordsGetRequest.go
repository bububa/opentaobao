package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取热词 APIRequest
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

func NewTaobaoAilabAicloudTopHotwordsGetRequest() *TaobaoAilabAicloudTopHotwordsGetRequest{
    return &TaobaoAilabAicloudTopHotwordsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.hotwords.get"
}

func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetBizClass(bizClass string) error {
    r.bizClass = bizClass
    r.Set("biz_class", bizClass)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetBizClass() string {
    return r.bizClass
}

func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetSchema() string {
    return r.schema
}

