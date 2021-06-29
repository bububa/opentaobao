package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) API请求
alibaba.wdk.ums.outbound

出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
*/
type AlibabaWdkUmsOutboundRequest struct {
    model.Params
    // 出库-ERP下发单请求dto
    _erpOutputOrderDto   *ErpOutputOrderDTO
}

// 初始化AlibabaWdkUmsOutboundRequest对象
func NewAlibabaWdkUmsOutboundRequest() *AlibabaWdkUmsOutboundRequest{
    return &AlibabaWdkUmsOutboundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOutboundRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.outbound"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOutboundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ErpOutputOrderDto Setter
// 出库-ERP下发单请求dto
func (r *AlibabaWdkUmsOutboundRequest) SetErpOutputOrderDto(_erpOutputOrderDto *ErpOutputOrderDTO) error {
    r._erpOutputOrderDto = _erpOutputOrderDto
    r.Set("erp_output_order_dto", _erpOutputOrderDto)
    return nil
}

// ErpOutputOrderDto Getter
func (r AlibabaWdkUmsOutboundRequest) GetErpOutputOrderDto() *ErpOutputOrderDTO {
    return r._erpOutputOrderDto
}
