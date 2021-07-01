package rhino

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步成衣仓盘点数据 API返回值 
taobao.rhino.supplychain.clothing.adjust

同步成衣仓盘点数据
*/
type TaobaoRhinoSupplychainClothingAdjustAPIResponse struct {
    model.CommonResponse
    TaobaoRhinoSupplychainClothingAdjustAPIResponseModel
}

// 同步成衣仓盘点数据 成功返回结果
type TaobaoRhinoSupplychainClothingAdjustAPIResponseModel struct {
    XMLName xml.Name `xml:"rhino_supplychain_clothing_adjust_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // code
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
