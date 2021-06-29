package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】火车票订单查询 APIResponse
alitrip.btrip.supplychain.train.search

【商旅】火车票订单查询
*/
type AlitripBtripSupplychainTrainSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainTrainSearchResponse
}

type AlitripBtripSupplychainTrainSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口对外数据透出
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
