package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) APIRequest
alibaba.wdk.ums.outbound

出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
*/
type AlibabaWdkUmsOutboundRequest struct {
    model.Params

    // 出库-ERP下发单请求dto
    erpOutputOrderDto   *ErpOutputOrderDto 

}

func NewAlibabaWdkUmsOutboundRequest() *AlibabaWdkUmsOutboundRequest{
    return &AlibabaWdkUmsOutboundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsOutboundRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.outbound"
}

func (r AlibabaWdkUmsOutboundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsOutboundRequest) SetErpOutputOrderDto(erpOutputOrderDto *ErpOutputOrderDto) error {
    r.erpOutputOrderDto = erpOutputOrderDto
    r.Set("erp_output_order_dto", erpOutputOrderDto)
    return nil
}

func (r AlibabaWdkUmsOutboundRequest) GetErpOutputOrderDto() *ErpOutputOrderDto {
    return r.erpOutputOrderDto
}

