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
    _userId   string
    // 业务类型
    _bizClass   string
    // schemeKey
    _schema   string
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
func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetUserId() string {
    return r._userId
}
// BizClass Setter
// 业务类型
func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetBizClass(_bizClass string) error {
    r._bizClass = _bizClass
    r.Set("biz_class", _bizClass)
    return nil
}

// BizClass Getter
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetBizClass() string {
    return r._bizClass
}
// Schema Setter
// schemeKey
func (r *TaobaoAilabAicloudTopHotwordsGetRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopHotwordsGetRequest) GetSchema() string {
    return r._schema
}
