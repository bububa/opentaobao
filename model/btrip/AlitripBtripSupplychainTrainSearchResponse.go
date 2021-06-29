package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】火车票订单查询 API返回值 
alitrip.btrip.supplychain.train.search

【商旅】火车票订单查询
*/
type AlitripBtripSupplychainTrainSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainTrainSearchResponse
}

// 【商旅】火车票订单查询 成功返回结果
type AlitripBtripSupplychainTrainSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口对外数据透出
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
