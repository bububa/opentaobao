package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条更新 API请求
taobao.alitrip.it.policy.update

销售规则更新接口，可以根据taobaoId或outId修改，outId不唯一时，不能用outId修改。
*/
type TaobaoAlitripItPolicyUpdateRequest struct {
    model.Params
    // 扩展字段
    _extendAttributes   string
    // 接入方产品id
    _outId   string
    // 淘宝政策id
    _taobaoId   int64
    // 国际机票销售规则
    _topPolicyDo   *TopPolicyDo
}

// 初始化TaobaoAlitripItPolicyUpdateRequest对象
func NewTaobaoAlitripItPolicyUpdateRequest() *TaobaoAlitripItPolicyUpdateRequest{
    return &TaobaoAlitripItPolicyUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendAttributes Setter
// 扩展字段
func (r *TaobaoAlitripItPolicyUpdateRequest) SetExtendAttributes(_extendAttributes string) error {
    r._extendAttributes = _extendAttributes
    r.Set("extend_attributes", _extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItPolicyUpdateRequest) GetExtendAttributes() string {
    return r._extendAttributes
}
// OutId Setter
// 接入方产品id
func (r *TaobaoAlitripItPolicyUpdateRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoAlitripItPolicyUpdateRequest) GetOutId() string {
    return r._outId
}
// TaobaoId Setter
// 淘宝政策id
func (r *TaobaoAlitripItPolicyUpdateRequest) SetTaobaoId(_taobaoId int64) error {
    r._taobaoId = _taobaoId
    r.Set("taobao_id", _taobaoId)
    return nil
}

// TaobaoId Getter
func (r TaobaoAlitripItPolicyUpdateRequest) GetTaobaoId() int64 {
    return r._taobaoId
}
// TopPolicyDo Setter
// 国际机票销售规则
func (r *TaobaoAlitripItPolicyUpdateRequest) SetTopPolicyDo(_topPolicyDo *TopPolicyDo) error {
    r._topPolicyDo = _topPolicyDo
    r.Set("top_policy_do", _topPolicyDo)
    return nil
}

// TopPolicyDo Getter
func (r TaobaoAlitripItPolicyUpdateRequest) GetTopPolicyDo() *TopPolicyDo {
    return r._topPolicyDo
}
