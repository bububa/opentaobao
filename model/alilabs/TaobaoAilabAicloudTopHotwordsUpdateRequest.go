package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新热词 API请求
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

// 初始化TaobaoAilabAicloudTopHotwordsUpdateRequest对象
func NewTaobaoAilabAicloudTopHotwordsUpdateRequest() *TaobaoAilabAicloudTopHotwordsUpdateRequest{
    return &TaobaoAilabAicloudTopHotwordsUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.hotwords.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schemaKey
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 三方用户id
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetUserId() string {
    return r.userId
}
// BizClass Setter
// 业务类型
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetBizClass(bizClass string) error {
    r.bizClass = bizClass
    r.Set("biz_class", bizClass)
    return nil
}

// BizClass Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetBizClass() string {
    return r.bizClass
}
// OpType Setter
// 操作类型
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetOpType(opType string) error {
    r.opType = opType
    r.Set("op_type", opType)
    return nil
}

// OpType Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetOpType() string {
    return r.opType
}
// Content Setter
// 热词内容
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetContent(content *HotWordsContent) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetContent() *HotWordsContent {
    return r.content
}
