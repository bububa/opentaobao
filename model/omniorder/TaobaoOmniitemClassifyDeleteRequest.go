package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一个分类 APIRequest
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

func NewTaobaoOmniitemClassifyDeleteRequest() *TaobaoOmniitemClassifyDeleteRequest{
    return &TaobaoOmniitemClassifyDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemClassifyDeleteRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.delete"
}

func (r TaobaoOmniitemClassifyDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemClassifyDeleteRequest) SetClassifyId(classifyId int64) error {
    r.classifyId = classifyId
    r.Set("classify_id", classifyId)
    return nil
}

func (r TaobaoOmniitemClassifyDeleteRequest) GetClassifyId() int64 {
    return r.classifyId
}

func (r *TaobaoOmniitemClassifyDeleteRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoOmniitemClassifyDeleteRequest) GetOperator() string {
    return r.operator
}

