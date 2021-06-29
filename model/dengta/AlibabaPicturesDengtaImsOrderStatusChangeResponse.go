package dengta

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天下秀回传订单执行状态变动 APIResponse
alibaba.pictures.dengta.ims.order.status.change

天下秀回传订单执行状态变动
*/
type AlibabaPicturesDengtaImsOrderStatusChangeAPIResponse struct {
    model.CommonResponse
    AlibabaPicturesDengtaImsOrderStatusChangeResponse
}

type AlibabaPicturesDengtaImsOrderStatusChangeResponse struct {
    XMLName xml.Name `xml:"alibaba_pictures_dengta_ims_order_status_change_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口出参
    
    Result   *GeneralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
