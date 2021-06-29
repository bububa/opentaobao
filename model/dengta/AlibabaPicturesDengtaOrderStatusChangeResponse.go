package dengta

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天下秀订单状态变更通知 APIResponse
alibaba.pictures.dengta.order.status.change

天下秀订单状态变更通知
*/
type AlibabaPicturesDengtaOrderStatusChangeAPIResponse struct {
    model.CommonResponse
    AlibabaPicturesDengtaOrderStatusChangeResponse
}

type AlibabaPicturesDengtaOrderStatusChangeResponse struct {
    XMLName xml.Name `xml:"alibaba_pictures_dengta_order_status_change_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
