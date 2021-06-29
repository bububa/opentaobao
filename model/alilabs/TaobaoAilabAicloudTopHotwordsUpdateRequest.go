package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新热词 APIRequest
taobao.ailab.aicloud.top.hotwords.update

更新ASR热词
*/
type TaobaoAilabAicloudTopHotwordsUpdateRequest struct {
    model.Params

    // schemaKey
    schema   string 

    // 三方用户id
    userId   string 

    // 业务类型
    bizClass   string 

    // 操作类型
    opType   string 

    // 热词内容
    content   *HotWordsContent 

}

func NewTaobaoAilabAicloudTopHotwordsUpdateRequest() *TaobaoAilabAicloudTopHotwordsUpdateRequest{
    return &TaobaoAilabAicloudTopHotwordsUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.hotwords.update"
}

func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetBizClass(bizClass string) error {
    r.bizClass = bizClass
    r.Set("biz_class", bizClass)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetBizClass() string {
    return r.bizClass
}

func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetOpType(opType string) error {
    r.opType = opType
    r.Set("op_type", opType)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetOpType() string {
    return r.opType
}

func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetContent(content *HotWordsContent) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetContent() *HotWordsContent {
    return r.content
}

