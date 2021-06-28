package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-店铺搜索 APIResponse
taobao.tbk.shop.get

淘宝客店铺查询
*/
type TaobaoTbkShopGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkShopGetResponse
}

type TaobaoTbkShopGetResponse struct {
    XMLName xml.Name `xml:"tbk_shop_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 淘宝客店铺
    
    Results   []NTbkShop `json:"results,omitempty" xml:"results>n_tbk_shop,omitempty"`
    
    
    // 搜索到符合条件的结果总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
