package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单消息状态回传 API返回值 
taobao.rdc.aligenius.ordermsg.update

用于订单消息处理状态回传
*/
type TaobaoRdcAligeniusOrdermsgUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusOrdermsgUpdateAPIResponseModel
}

// 订单消息状态回传 成功返回结果
type TaobaoRdcAligeniusOrdermsgUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"rdc_aligenius_ordermsg_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoRdcAligeniusOrdermsgUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
