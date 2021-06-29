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
type TaobaoAlitripItPolicyBatchdeleteRequest struct {
    model.Params
    // 航司二字码，完整匹配
    airline   string
    // 到达，，完整匹配
    arrCity   string
    // 舱位，，完整匹配
    cabin   string
    // 出发，，完整匹配
    depCity   string
    // 产品id，，完整匹配
    policyId   string
    // 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
    statusList   []int64
}

// 初始化TaobaoAlitripItPolicyBatchdeleteRequest对象
func NewTaobaoAlitripItPolicyBatchdeleteRequest() *TaobaoAlitripItPolicyBatchdeleteRequest{
    return &TaobaoAlitripItPolicyBatchdeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.batchdelete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Airline Setter
// 航司二字码，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteRequest) SetAirline(airline string) error {
    r.airline = airline
    r.Set("airline", airline)
    return nil
}

// Airline Getter
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetAirline() string {
    return r.airline
}
// ArrCity Setter
// 到达，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteRequest) SetArrCity(arrCity string) error {
    r.arrCity = arrCity
    r.Set("arr_city", arrCity)
    return nil
}

// ArrCity Getter
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetArrCity() string {
    return r.arrCity
}
// Cabin Setter
// 舱位，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteRequest) SetCabin(cabin string) error {
    r.cabin = cabin
    r.Set("cabin", cabin)
    return nil
}

// Cabin Getter
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetCabin() string {
    return r.cabin
}
// DepCity Setter
// 出发，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteRequest) SetDepCity(depCity string) error {
    r.depCity = depCity
    r.Set("dep_city", depCity)
    return nil
}

// DepCity Getter
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetDepCity() string {
    return r.depCity
}
// PolicyId Setter
// 产品id，，完整匹配
func (r *TaobaoAlitripItPolicyBatchdeleteRequest) SetPolicyId(policyId string) error {
    r.policyId = policyId
    r.Set("policy_id", policyId)
    return nil
}

// PolicyId Getter
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetPolicyId() string {
    return r.policyId
}
// StatusList Setter
// 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
func (r *TaobaoAlitripItPolicyBatchdeleteRequest) SetStatusList(statusList []int64) error {
    r.statusList = statusList
    r.Set("status_list", statusList)
    return nil
}

// StatusList Getter
func (r TaobaoAlitripItPolicyBatchdeleteRequest) GetStatusList() []int64 {
    return r.statusList
}
