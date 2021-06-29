package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康020接单 API返回值 
taobao.trade.drug.confirmorder

阿里健康020接单
*/
type TaobaoTradeDrugConfirmorderAPIResponse struct {
    model.CommonResponse
    TaobaoTradeDrugConfirmorderResponse
}

// 阿里健康020接单 成功返回结果
type TaobaoTradeDrugConfirmorderResponse struct {
    XMLName xml.Name `xml:"trade_drug_confirmorder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // life的返回值
    Result   *LifeResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
