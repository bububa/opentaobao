package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenSupplychainTrainTradeAPIResponse 商旅火车票交易流水接口 API返回值
// alitrip.btrip.open.supplychain.train.trade
//
// 商旅火车票交易流水接口
type AlitripBtripOpenSupplychainTrainTradeAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenSupplychainTrainTradeAPIResponseModel
}

// AlitripBtripOpenSupplychainTrainTradeAPIResponseModel is 商旅火车票交易流水接口 成功返回结果
type AlitripBtripOpenSupplychainTrainTradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_train_trade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
