package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店扫脸信息上传 APIResponse
taobao.xhotel.order.future.facescan.put

未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接
*/
type TaobaoXhotelOrderFutureFacescanPutAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderFutureFacescanPutResponse
}

type TaobaoXhotelOrderFutureFacescanPutResponse struct {
    XMLName xml.Name `xml:"xhotel_order_future_facescan_put_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoXhotelOrderFutureFacescanPutResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
