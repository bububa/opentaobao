package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychaintraindetailsearchvtwoAPIResponse 【商旅】火车票订单详情查询 API返回值
// alitrip.btrip.supplychain.train.detail.search.vtwo
//
// 【商旅】火车票订单详情查询
type AlitripbtripsupplychaintraindetailsearchvtwoAPIResponse struct {
	model.CommonResponse
	AlitripbtripsupplychaintraindetailsearchvtwoAPIResponseModel
}

// AlitripbtripsupplychaintraindetailsearchvtwoAPIResponseModel is 【商旅】火车票订单详情查询 成功返回结果
type AlitripbtripsupplychaintraindetailsearchvtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_detail_search_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
