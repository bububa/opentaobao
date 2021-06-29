package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
以盘点方式调整库存：传入商品实际库存 APIResponse
alibaba.mos.goods.synchinventorybycounting

以盘点方式调整库存：传入商品实际库存
盘点单自动判断数量增减
*/
type AlibabaMosGoodsSynchinventorybycountingAPIResponse struct {
    model.CommonResponse
    AlibabaMosGoodsSynchinventorybycountingResponse
}

type AlibabaMosGoodsSynchinventorybycountingResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_goods_synchinventorybycounting_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回盘点单号
    
    Result   *AlibabaMosGoodsSynchinventorybycountingResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
