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
    _schema   string
    // 三方用户id
    _userId   string
    // 业务类型
    _bizClass   string
    // 操作类型
    _opType   string
    // 热词内容
    _content   *HotWordsContent
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
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 三方用户id
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetUserId() string {
    return r._userId
}
// BizClass Setter
// 业务类型
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetBizClass(_bizClass string) error {
    r._bizClass = _bizClass
    r.Set("biz_class", _bizClass)
    return nil
}

// BizClass Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetBizClass() string {
    return r._bizClass
}
// OpType Setter
// 操作类型
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetOpType(_opType string) error {
    r._opType = _opType
    r.Set("op_type", _opType)
    return nil
}

// OpType Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetOpType() string {
    return r._opType
}
// Content Setter
// 热词内容
func (r *TaobaoAilabAicloudTopHotwordsUpdateRequest) SetContent(_content *HotWordsContent) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateRequest) GetContent() *HotWordsContent {
    return r._content
}
