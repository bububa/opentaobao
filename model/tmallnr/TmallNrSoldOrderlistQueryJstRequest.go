package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
已入塔商家查询订单列表 API请求
tmall.nr.sold.orderlist.query.jst

该服务用于已入聚石塔的商家，获取订单列表信息；
*/
type TmallNrSoldOrderlistQueryJstRequest struct {
    model.Params
    // 入参对象
    _param0   *NrTimingOrderSoldQueryReqDTO
}

// 初始化TmallNrSoldOrderlistQueryJstRequest对象
func NewTmallNrSoldOrderlistQueryJstRequest() *TmallNrSoldOrderlistQueryJstRequest{
    return &TmallNrSoldOrderlistQueryJstRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrSoldOrderlistQueryJstRequest) GetApiMethodName() string {
    return "tmall.nr.sold.orderlist.query.jst"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrSoldOrderlistQueryJstRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参对象
func (r *TmallNrSoldOrderlistQueryJstRequest) SetParam0(_param0 *NrTimingOrderSoldQueryReqDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallNrSoldOrderlistQueryJstRequest) GetParam0() *NrTimingOrderSoldQueryReqDTO {
    return r._param0
}
