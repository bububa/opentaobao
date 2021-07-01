package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】批量删除 API请求
taobao.alitrip.it.policy.batchdelete

批量删除销售规则，单次删除最大5万，大于5万时候提示失败，需要缩小删除条件。此接口同步返回任务id，异步执行操作。每个接入方最多同时只能有10个处理中的任务，超过后直接返回失败。
*/
type TaobaoAlitripItPolicyBatchdeleteAPIRequest struct {
    model.Params
    // 航司二字码，完整匹配
    _airline   string
    // 到达，，完整匹配
    _arrCity   string
    // 舱位，，完整匹配
    _cabin   string
    // 出发，，完整匹配
    _depCity   string
    // 产品id，，完整匹配
    _policyId   string
    // 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
    _statusList   []int64
}

// 初始化TaobaoAlitripItPolicyBatchdeleteAPIRequest对象
func NewTaobaoAlitripItPolicyBatchdeleteRequest() *TaobaoAlitripItPolicyBatchdeleteAPIRequest{
    return &TaobaoAlitripItPolicyBatchdeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.batchdelete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Airline Setter
// 航司二字码，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteAPIRequest) SetAirline(_airline string) error {
    r._airline = _airline
    r.Set("airline", _airline)
    return nil
}

// Airline Getter
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetAirline() string {
    return r._airline
}
// ArrCity Setter
// 到达，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteAPIRequest) SetArrCity(_arrCity string) error {
    r._arrCity = _arrCity
    r.Set("arr_city", _arrCity)
    return nil
}

// ArrCity Getter
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetArrCity() string {
    return r._arrCity
}
// Cabin Setter
// 舱位，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteAPIRequest) SetCabin(_cabin string) error {
    r._cabin = _cabin
    r.Set("cabin", _cabin)
    return nil
}

// Cabin Getter
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetCabin() string {
    return r._cabin
}
// DepCity Setter
// 出发，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteAPIRequest) SetDepCity(_depCity string) error {
    r._depCity = _depCity
    r.Set("dep_city", _depCity)
    return nil
}

// DepCity Getter
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetDepCity() string {
    return r._depCity
}
// PolicyId Setter
// 产品id，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteAPIRequest) SetPolicyId(_policyId string) error {
    r._policyId = _policyId
    r.Set("policy_id", _policyId)
    return nil
}

// PolicyId Getter
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetPolicyId() string {
    return r._policyId
}
// StatusList Setter
// 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
func (r *TaobaoAlitripItPolicyBatchdeleteAPIRequest) SetStatusList(_statusList []int64) error {
    r._statusList = _statusList
    r.Set("status_list", _statusList)
    return nil
}

// StatusList Getter
func (r TaobaoAlitripItPolicyBatchdeleteAPIRequest) GetStatusList() []int64 {
    return r._statusList
}
