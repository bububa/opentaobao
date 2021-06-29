package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺商品接口 APIResponse
taobao.txp.item.itemlistget

淘小铺商品的查询服务。
*/
type TaobaoTxpItemItemlistgetAPIResponse struct {
    model.CommonResponse
    TaobaoTxpItemItemlistgetResponse
}

type TaobaoTxpItemItemlistgetResponse struct {
    XMLName xml.Name `xml:"txp_item_itemlistget_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *TaobaoTxpItemItemlistgetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
