package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-淘宝客商品详情查询(简版) APIResponse
taobao.tbk.item.info.get

淘宝客商品详情查询（简版）
*/
type TaobaoTbkItemInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkItemInfoGetResponse
}

type TaobaoTbkItemInfoGetResponse struct {
    XMLName xml.Name `xml:"tbk_item_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 淘宝客商品
    
    Results   []NTbkItem `json:"results,omitempty" xml:"results>n_tbk_item,omitempty"`
    
    
}
