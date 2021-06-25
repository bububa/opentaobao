package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售商品发货 APIRequest
alibaba.mj.oc.outbound

用于接收发货的数据
*/
type AlibabaMjOcOutboundRequest struct {
    model.Params

    // 发货信息
    goodsOutbound   *GoodsOutboundDTO 

}

func NewAlibabaMjOcOutboundRequest() *AlibabaMjOcOutboundRequest{
    return &AlibabaMjOcOutboundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcOutboundRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.outbound"
}

func (r AlibabaMjOcOutboundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcOutboundRequest) SetGoodsOutbound(goodsOutbound *GoodsOutboundDTO) error {
    r.goodsOutbound = goodsOutbound
    r.Set("goods_outbound", goodsOutbound)
    return nil
}

func (r AlibabaMjOcOutboundRequest) GetGoodsOutbound() *GoodsOutboundDTO {
    return r.goodsOutbound
}

