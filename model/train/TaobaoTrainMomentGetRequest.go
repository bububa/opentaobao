package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票时刻表 APIRequest
taobao.train.moment.get

查询火车票车次时刻表
*/
type TaobaoTrainMomentGetRequest struct {
    model.Params

    // 出参
    param   *TrainMomentTopParam 

}

func NewTaobaoTrainMomentGetRequest() *TaobaoTrainMomentGetRequest{
    return &TaobaoTrainMomentGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainMomentGetRequest) GetApiMethodName() string {
    return "taobao.train.moment.get"
}

func (r TaobaoTrainMomentGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainMomentGetRequest) SetParam(param *TrainMomentTopParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoTrainMomentGetRequest) GetParam() *TrainMomentTopParam {
    return r.param
}

