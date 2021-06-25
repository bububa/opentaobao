package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库-ERP下发单 APIRequest
alibaba.wdk.ums.inbound

入库-ERP下发单
*/
type AlibabaWdkUmsInboundRequest struct {
    model.Params

    // 1
    erpArrivalnoticeDto   *ErpArrivalNoticeDto 

}

func NewAlibabaWdkUmsInboundRequest() *AlibabaWdkUmsInboundRequest{
    return &AlibabaWdkUmsInboundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsInboundRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inbound"
}

func (r AlibabaWdkUmsInboundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsInboundRequest) SetErpArrivalnoticeDto(erpArrivalnoticeDto *ErpArrivalNoticeDto) error {
    r.erpArrivalnoticeDto = erpArrivalnoticeDto
    r.Set("erp_arrivalnotice_dto", erpArrivalnoticeDto)
    return nil
}

func (r AlibabaWdkUmsInboundRequest) GetErpArrivalnoticeDto() *ErpArrivalNoticeDto {
    return r.erpArrivalnoticeDto
}

