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
    _fromCase   string
    // 我的家房间json
    _toCase   string
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
func (r *TaobaoHomeaiAlgPredictRequest) SetFromCase(_fromCase string) error {
    r._fromCase = _fromCase
    r.Set("from_case", _fromCase)
    return nil
}

// FromCase Getter
func (r TaobaoHomeaiAlgPredictRequest) GetFromCase() string {
    return r._fromCase
}
// ToCase Setter
// 我的家房间json
func (r *TaobaoHomeaiAlgPredictRequest) SetToCase(_toCase string) error {
    r._toCase = _toCase
    r.Set("to_case", _toCase)
    return nil
}

// ToCase Getter
func (r TaobaoHomeaiAlgPredictRequest) GetToCase() string {
    return r._toCase
}
