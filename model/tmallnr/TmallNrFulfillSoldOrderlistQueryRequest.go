package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售商获取品牌商的特定订单列表 APIRequest
tmall.nr.fulfill.sold.orderlist.query

零售商获取品牌商的特定订单列表，后端服务有零售商和品牌商的绑定关系，存在开关控制；同时后端存在定时送业务等特殊业务的校验，非同城配送业务不能返回，返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
*/
type TmallNrFulfillSoldOrderlistQueryRequest struct {
    model.Params

    // 入参对象
    param0   *NrTimingOrderSoldQueryReqDto 

}

func NewTmallNrFulfillSoldOrderlistQueryRequest() *TmallNrFulfillSoldOrderlistQueryRequest{
    return &TmallNrFulfillSoldOrderlistQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrFulfillSoldOrderlistQueryRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.sold.orderlist.query"
}

func (r TmallNrFulfillSoldOrderlistQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrFulfillSoldOrderlistQueryRequest) SetParam0(param0 *NrTimingOrderSoldQueryReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallNrFulfillSoldOrderlistQueryRequest) GetParam0() *NrTimingOrderSoldQueryReqDto {
    return r.param0
}

