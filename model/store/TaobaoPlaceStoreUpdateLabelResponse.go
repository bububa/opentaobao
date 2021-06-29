package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店标签更新接口 API返回值 
taobao.place.store.update.label

更新商户门店标签（服务、权益、标签）接口
*/
type TaobaoPlaceStoreUpdateLabelAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreUpdateLabelResponse
}

// 商户门店标签更新接口 成功返回结果
type TaobaoPlaceStoreUpdateLabelResponse struct {
    XMLName xml.Name `xml:"place_store_update_label_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *UpdateResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
