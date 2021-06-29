package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单消息状态回传 APIResponse
taobao.rdc.aligenius.ordermsg.update

用于订单消息处理状态回传
*/
type TaobaoRdcAligeniusOrdermsgUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusOrdermsgUpdateResponse
}

type TaobaoRdcAligeniusOrdermsgUpdateResponse struct {
    XMLName xml.Name `xml:"rdc_aligenius_ordermsg_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoRdcAligeniusOrdermsgUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
