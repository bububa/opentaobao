package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康020拒单 APIResponse
taobao.trade.drug.refuseorder

阿里健康020拒单
*/
type TaobaoTradeDrugRefuseorderAPIResponse struct {
    model.CommonResponse
    TaobaoTradeDrugRefuseorderResponse
}

type TaobaoTradeDrugRefuseorderResponse struct {
    XMLName xml.Name `xml:"trade_drug_refuseorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // life的返回值
    
    Result   *LifeResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
