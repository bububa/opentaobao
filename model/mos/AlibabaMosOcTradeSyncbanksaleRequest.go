package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云闪付、银行卡销售数据回传接口 APIRequest
alibaba.mos.oc.trade.syncbanksale

云闪付、银行卡销售数据回传
*/
type AlibabaMosOcTradeSyncbanksaleRequest struct {
    model.Params

    // pos云闪付、银行卡销售数据
    posBankSaleInfoDto   *PosBankSaleInfoDto 

}

func NewAlibabaMosOcTradeSyncbanksaleRequest() *AlibabaMosOcTradeSyncbanksaleRequest{
    return &AlibabaMosOcTradeSyncbanksaleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOcTradeSyncbanksaleRequest) GetApiMethodName() string {
    return "alibaba.mos.oc.trade.syncbanksale"
}

func (r AlibabaMosOcTradeSyncbanksaleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOcTradeSyncbanksaleRequest) SetPosBankSaleInfoDto(posBankSaleInfoDto *PosBankSaleInfoDto) error {
    r.posBankSaleInfoDto = posBankSaleInfoDto
    r.Set("pos_bank_sale_info_dto", posBankSaleInfoDto)
    return nil
}

func (r AlibabaMosOcTradeSyncbanksaleRequest) GetPosBankSaleInfoDto() *PosBankSaleInfoDto {
    return r.posBankSaleInfoDto
}

