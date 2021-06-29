package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝退货 API返回值 
taobao.rp.returngoods.refuse

卖家拒绝退货，目前仅支持天猫退货。
*/
type TaobaoRpReturngoodsRefuseAPIResponse struct {
    model.CommonResponse
    TaobaoRpReturngoodsRefuseResponse
}

// 卖家拒绝退货 成功返回结果
type TaobaoRpReturngoodsRefuseResponse struct {
    XMLName xml.Name `xml:"rp_returngoods_refuse_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // asdf
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
