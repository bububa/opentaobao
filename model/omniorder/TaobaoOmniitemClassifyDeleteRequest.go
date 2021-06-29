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
    classifyId   int64
    // 操作人信息（暂时不填）
    operator   string
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
func (r *TaobaoOmniitemClassifyDeleteRequest) SetClassifyId(classifyId int64) error {
    r.classifyId = classifyId
    r.Set("classify_id", classifyId)
    return nil
}

// ClassifyId Getter
func (r TaobaoOmniitemClassifyDeleteRequest) GetClassifyId() int64 {
    return r.classifyId
}
// Operator Setter
// 操作人信息（暂时不填）
func (r *TaobaoOmniitemClassifyDeleteRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoOmniitemClassifyDeleteRequest) GetOperator() string {
    return r.operator
}
