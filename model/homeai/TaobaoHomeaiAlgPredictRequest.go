package homeai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
硬装预测接口 API请求
taobao.homeai.alg.predict

居然之家硬装预测服务
*/
type TaobaoHomeaiAlgPredictRequest struct {
    model.Params
    // 来源房间json
    fromCase   string
    // 我的家房间json
    toCase   string
}

// 初始化TaobaoHomeaiAlgPredictRequest对象
func NewTaobaoHomeaiAlgPredictRequest() *TaobaoHomeaiAlgPredictRequest{
    return &TaobaoHomeaiAlgPredictRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoHomeaiAlgPredictRequest) GetApiMethodName() string {
    return "taobao.homeai.alg.predict"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoHomeaiAlgPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FromCase Setter
// 来源房间json
func (r *TaobaoHomeaiAlgPredictRequest) SetFromCase(fromCase string) error {
    r.fromCase = fromCase
    r.Set("from_case", fromCase)
    return nil
}

// FromCase Getter
func (r TaobaoHomeaiAlgPredictRequest) GetFromCase() string {
    return r.fromCase
}
// ToCase Setter
// 我的家房间json
func (r *TaobaoHomeaiAlgPredictRequest) SetToCase(toCase string) error {
    r.toCase = toCase
    r.Set("to_case", toCase)
    return nil
}

// ToCase Getter
func (r TaobaoHomeaiAlgPredictRequest) GetToCase() string {
    return r.toCase
}
