package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条查询 API请求
taobao.alitrip.it.policy.get

通过此接口可以查询单条销售规则的详情，可以根据taobaoId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据。taobaoId为新增成功时候返回的唯一id，outId为新增时的policy_id（产品编号）
*/
type TaobaoAlitripItPolicyGetRequest struct {
    model.Params
    // 预留扩展字段
    _extendAttributes   string
    // 接入方产品编号
    _outId   string
    // 淘宝政策id
    _taobaoId   int64
}

// 初始化TaobaoAlitripItPolicyGetRequest对象
func NewTaobaoAlitripItPolicyGetRequest() *TaobaoAlitripItPolicyGetRequest{
    return &TaobaoAlitripItPolicyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendAttributes Setter
// 预留扩展字段
func (r *TaobaoAlitripItPolicyGetRequest) SetExtendAttributes(_extendAttributes string) error {
    r._extendAttributes = _extendAttributes
    r.Set("extend_attributes", _extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItPolicyGetRequest) GetExtendAttributes() string {
    return r._extendAttributes
}
// OutId Setter
// 接入方产品编号
func (r *TaobaoAlitripItPolicyGetRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoAlitripItPolicyGetRequest) GetOutId() string {
    return r._outId
}
// TaobaoId Setter
// 淘宝政策id
func (r *TaobaoAlitripItPolicyGetRequest) SetTaobaoId(_taobaoId int64) error {
    r._taobaoId = _taobaoId
    r.Set("taobao_id", _taobaoId)
    return nil
}

// TaobaoId Getter
func (r TaobaoAlitripItPolicyGetRequest) GetTaobaoId() int64 {
    return r._taobaoId
}
