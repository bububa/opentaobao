package homeai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
硬装预测接口 APIRequest
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

func NewTaobaoHomeaiAlgPredictRequest() *TaobaoHomeaiAlgPredictRequest{
    return &TaobaoHomeaiAlgPredictRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoHomeaiAlgPredictRequest) GetApiMethodName() string {
    return "taobao.homeai.alg.predict"
}

func (r TaobaoHomeaiAlgPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoHomeaiAlgPredictRequest) SetFromCase(fromCase string) error {
    r.fromCase = fromCase
    r.Set("from_case", fromCase)
    return nil
}

func (r TaobaoHomeaiAlgPredictRequest) GetFromCase() string {
    return r.fromCase
}

func (r *TaobaoHomeaiAlgPredictRequest) SetToCase(toCase string) error {
    r.toCase = toCase
    r.Set("to_case", toCase)
    return nil
}

func (r TaobaoHomeaiAlgPredictRequest) GetToCase() string {
    return r.toCase
}

