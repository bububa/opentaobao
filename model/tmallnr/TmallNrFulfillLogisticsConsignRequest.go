package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城配门店备货通知 APIRequest
tmall.nr.fulfill.logistics.consign

同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了；
*/
type TmallNrFulfillLogisticsConsignRequest struct {
    model.Params

    // 入参对象
    param0   *NrStoreGoodsReadyReqDto 

}

func NewTmallNrFulfillLogisticsConsignRequest() *TmallNrFulfillLogisticsConsignRequest{
    return &TmallNrFulfillLogisticsConsignRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrFulfillLogisticsConsignRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.logistics.consign"
}

func (r TmallNrFulfillLogisticsConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrFulfillLogisticsConsignRequest) SetParam0(param0 *NrStoreGoodsReadyReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallNrFulfillLogisticsConsignRequest) GetParam0() *NrStoreGoodsReadyReqDto {
    return r.param0
}

