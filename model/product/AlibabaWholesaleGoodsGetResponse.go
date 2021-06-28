package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询阿里巴巴批发市场商品详情 APIResponse
alibaba.wholesale.goods.get

查询阿里巴巴批发市场商品详情
*/
type AlibabaWholesaleGoodsGetAPIResponse struct {
    model.CommonResponse
    AlibabaWholesaleGoodsGetResponse
}

type AlibabaWholesaleGoodsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wholesale_goods_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // wholesale goods detail result
    
    WholesaleGoodsResult   *WholesaleGoodsOpenResult `json:"wholesale_goods_result,omitempty" xml:"wholesale_goods_result,omitempty"`

    
}
