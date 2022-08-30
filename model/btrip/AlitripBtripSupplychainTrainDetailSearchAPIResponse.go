package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainDetailSearchAPIResponse 【商旅】火车票订单详情查询 API返回值
// alitrip.btrip.supplychain.train.detail.search
//
// 【商旅】火车票订单详情查询
type AlitripBtripSupplychainTrainDetailSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainTrainDetailSearchAPIResponseModel
}

// AlitripBtripSupplychainTrainDetailSearchAPIResponseModel is 【商旅】火车票订单详情查询 成功返回结果
type AlitripBtripSupplychainTrainDetailSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_detail_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
