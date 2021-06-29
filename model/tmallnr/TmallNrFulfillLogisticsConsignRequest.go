package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城配门店备货通知 API请求
tmall.nr.fulfill.logistics.consign

同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了；
*/
type TmallNrFulfillLogisticsConsignRequest struct {
    model.Params
    // 入参对象
    _param0   *NrStoreGoodsReadyReqDto
}

// 初始化TmallNrFulfillLogisticsConsignRequest对象
func NewTmallNrFulfillLogisticsConsignRequest() *TmallNrFulfillLogisticsConsignRequest{
    return &TmallNrFulfillLogisticsConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrFulfillLogisticsConsignRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.logistics.consign"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrFulfillLogisticsConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参对象
func (r *TmallNrFulfillLogisticsConsignRequest) SetParam0(_param0 *NrStoreGoodsReadyReqDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallNrFulfillLogisticsConsignRequest) GetParam0() *NrStoreGoodsReadyReqDto {
    return r._param0
}
