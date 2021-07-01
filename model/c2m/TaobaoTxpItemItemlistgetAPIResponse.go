package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺商品接口 API返回值 
taobao.txp.item.itemlistget

淘小铺商品的查询服务。
*/
type TaobaoTxpItemItemlistgetAPIResponse struct {
    model.CommonResponse
    TaobaoTxpItemItemlistgetAPIResponseModel
}

// 淘小铺商品接口 成功返回结果
type TaobaoTxpItemItemlistgetAPIResponseModel struct {
    XMLName xml.Name `xml:"txp_item_itemlistget_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *TaobaoTxpItemItemlistgetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
