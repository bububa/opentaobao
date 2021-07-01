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
type TaobaoAilabAicloudTopHotwordsUpdateAPIRequest struct {
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

// 初始化TaobaoAilabAicloudTopHotwordsUpdateAPIRequest对象
func NewTaobaoAilabAicloudTopHotwordsUpdateRequest() *TaobaoAilabAicloudTopHotwordsUpdateAPIRequest{
    return &TaobaoAilabAicloudTopHotwordsUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.hotwords.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schemaKey
func (r *TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 三方用户id
func (r *TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) GetUserId() string {
    return r._userId
}
// BizClass Setter
// 业务类型
func (r *TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) SetBizClass(_bizClass string) error {
    r._bizClass = _bizClass
    r.Set("biz_class", _bizClass)
    return nil
}

// BizClass Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) GetBizClass() string {
    return r._bizClass
}
// OpType Setter
// 操作类型
func (r *TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) SetOpType(_opType string) error {
    r._opType = _opType
    r.Set("op_type", _opType)
    return nil
}

// OpType Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) GetOpType() string {
    return r._opType
}
// Content Setter
// 热词内容
func (r *TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) SetContent(_content *HotWordsContent) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoAilabAicloudTopHotwordsUpdateAPIRequest) GetContent() *HotWordsContent {
    return r._content
}
