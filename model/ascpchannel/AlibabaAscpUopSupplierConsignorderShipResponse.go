package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
履约单商家仓发货结果回传服务 APIResponse
alibaba.ascp.uop.supplier.consignorder.ship

ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
*/
type AlibabaAscpUopSupplierConsignorderShipAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopSupplierConsignorderShipResponse
}

type AlibabaAscpUopSupplierConsignorderShipResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_consignorder_ship_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    ConsignorderShipResponse   *ResultWrapper `json:"consignorder_ship_response,omitempty" xml:"consignorder_ship_response,omitempty"`

    
}
