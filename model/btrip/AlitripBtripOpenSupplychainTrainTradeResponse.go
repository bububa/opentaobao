package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅火车票交易流水接口 APIResponse
alitrip.btrip.open.supplychain.train.trade

商旅火车票交易流水接口
*/
type AlitripBtripOpenSupplychainTrainTradeAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenSupplychainTrainTradeResponse
}

type AlitripBtripOpenSupplychainTrainTradeResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_train_trade_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
