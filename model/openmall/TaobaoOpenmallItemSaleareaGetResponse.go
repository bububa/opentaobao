package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品可售区域 API返回值 
taobao.openmall.item.salearea.get

获取openmall商品的可售区域
*/
type TaobaoOpenmallItemSaleareaGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallItemSaleareaGetResponse
}

// 查询商品可售区域 成功返回结果
type TaobaoOpenmallItemSaleareaGetResponse struct {
    XMLName xml.Name `xml:"openmall_item_salearea_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoOpenmallItemSaleareaGetResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
