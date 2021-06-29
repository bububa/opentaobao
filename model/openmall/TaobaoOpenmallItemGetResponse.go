package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品详情物料 API返回值 
taobao.openmall.item.get

获取联盟开放的openmall商品
*/
type TaobaoOpenmallItemGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallItemGetResponse
}

// 获取商品详情物料 成功返回结果
type TaobaoOpenmallItemGetResponse struct {
    XMLName xml.Name `xml:"openmall_item_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoOpenmallItemGetResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
