package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝退货 APIResponse
taobao.rp.returngoods.refuse

卖家拒绝退货，目前仅支持天猫退货。
*/
type TaobaoRpReturngoodsRefuseAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rp_returngoods_refuse_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // asdf
    
    Result   bool `json:"result,omitempty" xml:"