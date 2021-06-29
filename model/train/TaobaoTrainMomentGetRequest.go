package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票时刻表 API请求
taobao.train.moment.get

查询火车票车次时刻表
*/
type TaobaoTrainMomentGetRequest struct {
    model.Params
    // 出参
    param   *TrainMomentTopParam
}

// 初始化TaobaoTrainMomentGetRequest对象
func NewTaobaoTrainMomentGetRequest() *TaobaoTrainMomentGetRequest{
    return &TaobaoTrainMomentGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainMomentGetRequest) GetApiMethodName() string {
    return "taobao.train.moment.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainMomentGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 出参
func (r *TaobaoTrainMomentGetRequest) SetParam(param *TrainMomentTopParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r TaobaoTrainMomentGetRequest) GetParam() *TrainMomentTopParam {
    return r.param
}
