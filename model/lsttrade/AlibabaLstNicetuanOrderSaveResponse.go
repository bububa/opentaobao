package lsttrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
十荟团订单同步至零售通 APIResponse
alibaba.lst.nicetuan.order.save

十荟团订单同步至零售通，十荟团单向写到零售通
*/
type AlibabaLstNicetuanOrderSaveAPIResponse struct {
    model.CommonResponse
    AlibabaLstNicetuanOrderSaveResponse
}

type AlibabaLstNicetuanOrderSaveResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_nicetuan_order_save_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *HsfResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}
