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
type TaobaoTrainMomentGetAPIRequest struct {
    model.Params
    // 出参
    _param   *TrainMomentTopParam
}

// 初始化TaobaoTrainMomentGetAPIRequest对象
func NewTaobaoTrainMomentGetRequest() *TaobaoTrainMomentGetAPIRequest{
    return &TaobaoTrainMomentGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainMomentGetAPIRequest) GetApiMethodName() string {
    return "taobao.train.moment.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainMomentGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 出参
func (r *TaobaoTrainMomentGetAPIRequest) SetParam(_param *TrainMomentTopParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoTrainMomentGetAPIRequest) GetParam() *TrainMomentTopParam {
    return r._param
}
