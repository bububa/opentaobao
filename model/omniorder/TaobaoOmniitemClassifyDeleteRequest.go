package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一个分类 API请求
taobao.omniitem.classify.delete

删除一个分类
*/
type TaobaoOmniitemClassifyDeleteRequest struct {
    model.Params
    // 分类ID
    _classifyId   int64
    // 操作人信息（暂时不填）
    _operator   string
}

// 初始化TaobaoOmniitemClassifyDeleteRequest对象
func NewTaobaoOmniitemClassifyDeleteRequest() *TaobaoOmniitemClassifyDeleteRequest{
    return &TaobaoOmniitemClassifyDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyDeleteRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClassifyId Setter
// 分类ID
func (r *TaobaoOmniitemClassifyDeleteRequest) SetClassifyId(_classifyId int64) error {
    r._classifyId = _classifyId
    r.Set("classify_id", _classifyId)
    return nil
}

// ClassifyId Getter
func (r TaobaoOmniitemClassifyDeleteRequest) GetClassifyId() int64 {
    return r._classifyId
}
// Operator Setter
// 操作人信息（暂时不填）
func (r *TaobaoOmniitemClassifyDeleteRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoOmniitemClassifyDeleteRequest) GetOperator() string {
    return r._operator
}
