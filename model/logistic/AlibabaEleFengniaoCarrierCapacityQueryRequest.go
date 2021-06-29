package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按照门店查询骑手运力状态查询 API请求
alibaba.ele.fengniao.carrier.capacity.query

提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量
*/
type AlibabaEleFengniaoCarrierCapacityQueryRequest struct {
    model.Params
    // 系统自动生成
    _param   *Param
}

// 初始化AlibabaEleFengniaoCarrierCapacityQueryRequest对象
func NewAlibabaEleFengniaoCarrierCapacityQueryRequest() *AlibabaEleFengniaoCarrierCapacityQueryRequest{
    return &AlibabaEleFengniaoCarrierCapacityQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCarrierCapacityQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.carrier.capacity.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCarrierCapacityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoCarrierCapacityQueryRequest) SetParam(_param *Param) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoCarrierCapacityQueryRequest) GetParam() *Param {
    return r._param
}
