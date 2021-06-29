package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康020接单 APIResponse
taobao.trade.drug.confirmorder

阿里健康020接单
*/
type TaobaoTradeDrugConfirmorderAPIResponse struct {
    model.CommonResponse
    TaobaoTradeDrugConfirmorderResponse
}

type TaobaoTradeDrugConfirmorderResponse struct {
    XMLName xml.Name `xml:"trade_drug_confirmorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // life的返回值
    
    Result   *LifeResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
