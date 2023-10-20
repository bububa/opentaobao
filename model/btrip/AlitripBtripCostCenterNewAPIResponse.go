package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcenternewAPIResponse 新建外部成本中心 API返回值
// alitrip.btrip.cost.center.new
//
// 新建外部成本中心
type AlitripbtripcostcenternewAPIResponse struct {
	model.CommonResponse
	AlitripbtripcostcenternewAPIResponseModel
}

// AlitripbtripcostcenternewAPIResponseModel is 新建外部成本中心 成功返回结果
type AlitripbtripcostcenternewAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}
