package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtipcostcenterqueryAPIResponse 查询外部成本中心 API返回值
// alitrip.btip.cost.center.query
//
// 查询外部成本中心
type AlitripbtipcostcenterqueryAPIResponse struct {
	model.CommonResponse
	AlitripbtipcostcenterqueryAPIResponseModel
}

// AlitripbtipcostcenterqueryAPIResponseModel is 查询外部成本中心 成功返回结果
type AlitripbtipcostcenterqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btip_cost_center_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}
