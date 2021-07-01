package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云闪付、银行卡销售数据回传接口 API请求
alibaba.mos.oc.trade.syncbanksale

云闪付、银行卡销售数据回传
*/
type AlibabaMosOcTradeSyncbanksaleAPIRequest struct {
    model.Params
    // pos云闪付、银行卡销售数据
    _posBankSaleInfoDto   *PosBankSaleInfoDTO
}

// 初始化AlibabaMosOcTradeSyncbanksaleAPIRequest对象
func NewAlibabaMosOcTradeSyncbanksaleRequest() *AlibabaMosOcTradeSyncbanksaleAPIRequest{
    return &AlibabaMosOcTradeSyncbanksaleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOcTradeSyncbanksaleAPIRequest) GetApiMethodName() string {
    return "alibaba.mos.oc.trade.syncbanksale"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOcTradeSyncbanksaleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosBankSaleInfoDto Setter
// pos云闪付、银行卡销售数据
func (r *AlibabaMosOcTradeSyncbanksaleAPIRequest) SetPosBankSaleInfoDto(_posBankSaleInfoDto *PosBankSaleInfoDTO) error {
    r._posBankSaleInfoDto = _posBankSaleInfoDto
    r.Set("pos_bank_sale_info_dto", _posBankSaleInfoDto)
    return nil
}

// PosBankSaleInfoDto Getter
func (r AlibabaMosOcTradeSyncbanksaleAPIRequest) GetPosBankSaleInfoDto() *PosBankSaleInfoDTO {
    return r._posBankSaleInfoDto
}
