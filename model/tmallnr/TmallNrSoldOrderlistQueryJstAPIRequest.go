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
type TmallNrSoldOrderlistQueryJstAPIRequest struct {
    model.Params
    // 入参对象
    _param0   *NrTimingOrderSoldQueryReqDto
}

// 初始化TmallNrSoldOrderlistQueryJstAPIRequest对象
func NewTmallNrSoldOrderlistQueryJstRequest() *TmallNrSoldOrderlistQueryJstAPIRequest{
    return &TmallNrSoldOrderlistQueryJstAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrSoldOrderlistQueryJstAPIRequest) GetApiMethodName() string {
    return "tmall.nr.sold.orderlist.query.jst"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrSoldOrderlistQueryJstAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参对象
func (r *TmallNrSoldOrderlistQueryJstAPIRequest) SetParam0(_param0 *NrTimingOrderSoldQueryReqDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallNrSoldOrderlistQueryJstAPIRequest) GetParam0() *NrTimingOrderSoldQueryReqDto {
    return r._param0
}
