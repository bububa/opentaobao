package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
调整库存 APIResponse
alibaba.mos.goods.adjust

库存调整接口
*/
type AlibabaMosGoodsAdjustAPIResponse struct {
    model.CommonResponse
    AlibabaMosGoodsAdjustResponse
}

type AlibabaMosGoodsAdjustResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_goods_adjust_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 库存调整单号
    
    Result   *AlibabaMosGoodsAdjustResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
