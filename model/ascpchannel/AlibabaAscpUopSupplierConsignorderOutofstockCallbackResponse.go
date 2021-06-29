package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
履约单纬度的仓缺货回告服务 APIResponse
alibaba.ascp.uop.supplier.consignorder.outofstock.callback

商家仓履约单纬度的仓缺货回告接口
*/
type AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopSupplierConsignorderOutofstockCallbackResponse
}

type AlibabaAscpUopSupplierConsignorderOutofstockCallbackResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_consignorder_outofstock_callback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    ConsignorderOutofstockCallbackResponse   *ResultWrapper `json:"consignorder_outofstock_callback_response,omitempty" xml:"consignorder_outofstock_callback_response,omitempty"`

    
}
