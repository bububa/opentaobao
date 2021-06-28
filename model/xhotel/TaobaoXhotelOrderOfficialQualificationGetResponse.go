package xhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住用户资质校验 APIResponse
taobao.xhotel.order.official.qualification.get

官网信用住在下单前对用户进行资质校验，资质校验通过才能进行信用支付
*/
type TaobaoXhotelOrderOfficialQualificationGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"xhotel_order_official_qualification_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 阿里旅行中间调用URL
    
    TransferUrl   string `json:"transfer_url,omitempty" xml:"