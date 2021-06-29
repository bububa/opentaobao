package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条新增 API请求
taobao.alitrip.it.policy.add

销售规则新增，成功返回taobaoId
*/
type TaobaoAlitripItPolicyAddRequest struct {
    model.Params
    // 扩展字段
    _extendAttributes   string
    // 国际机票销售规则
    _topPolicyDo   *TopPolicyDo
}

// 初始化TaobaoAlitripItPolicyAddRequest对象
func NewTaobaoAlitripItPolicyAddRequest() *TaobaoAlitripItPolicyAddRequest{
    return &TaobaoAlitripItPolicyAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendAttributes Setter
// 扩展字段
func (r *TaobaoAlitripItPolicyAddRequest) SetExtendAttributes(_extendAttributes string) error {
    r._extendAttributes = _extendAttributes
    r.Set("extend_attributes", _extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItPolicyAddRequest) GetExtendAttributes() string {
    return r._extendAttributes
}
// TopPolicyDo Setter
// 国际机票销售规则
func (r *TaobaoAlitripItPolicyAddRequest) SetTopPolicyDo(_topPolicyDo *TopPolicyDo) error {
    r._topPolicyDo = _topPolicyDo
    r.Set("top_policy_do", _topPolicyDo)
    return nil
}

// TopPolicyDo Getter
func (r TaobaoAlitripItPolicyAddRequest) GetTopPolicyDo() *TopPolicyDo {
    return r._topPolicyDo
}
