package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机库存商品同步 API返回值 
alibaba.retail.vending.road.item.sync

贩卖机库存商品同步
*/
type AlibabaRetailVendingRoadItemSyncAPIResponse struct {
    model.CommonResponse
    AlibabaRetailVendingRoadItemSyncAPIResponseModel
}

// 贩卖机库存商品同步 成功返回结果
type AlibabaRetailVendingRoadItemSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_retail_vending_road_item_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 成功标识
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 是否成功
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`
    // 错误码
    ECode   string `json:"e_code,omitempty" xml:"e_code,omitempty"`
    // 错误消息
    EMsg   string `json:"e_msg,omitempty" xml:"e_msg,omitempty"`
}
