package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库-ERP下发单 API请求
alibaba.wdk.ums.inbound

入库-ERP下发单
*/
type AlibabaWdkUmsInboundRequest struct {
    model.Params
    // 1
    erpArrivalnoticeDto   *ErpArrivalNoticeDto
}

// 初始化AlibabaWdkUmsInboundRequest对象
func NewAlibabaWdkUmsInboundRequest() *AlibabaWdkUmsInboundRequest{
    return &AlibabaWdkUmsInboundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInboundRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inbound"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInboundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ErpArrivalnoticeDto Setter
// 1
func (r *AlibabaWdkUmsInboundRequest) SetErpArrivalnoticeDto(erpArrivalnoticeDto *ErpArrivalNoticeDto) error {
    r.erpArrivalnoticeDto = erpArrivalnoticeDto
    r.Set("erp_arrivalnotice_dto", erpArrivalnoticeDto)
    return nil
}

// ErpArrivalnoticeDto Getter
func (r AlibabaWdkUmsInboundRequest) GetErpArrivalnoticeDto() *ErpArrivalNoticeDto {
    return r.erpArrivalnoticeDto
}
