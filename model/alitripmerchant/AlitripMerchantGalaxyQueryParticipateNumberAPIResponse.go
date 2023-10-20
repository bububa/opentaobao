package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyqueryparticipatenumberAPIResponse 星河-抽奖活动次数查询 API返回值
// alitrip.merchant.galaxy.query.participate.number
//
// 雅高小程序抽奖活动，查询用户抽奖次数
type AlitripmerchantgalaxyqueryparticipatenumberAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyqueryparticipatenumberAPIResponseModel
}

// AlitripmerchantgalaxyqueryparticipatenumberAPIResponseModel is 星河-抽奖活动次数查询 成功返回结果
type AlitripmerchantgalaxyqueryparticipatenumberAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_query_participate_number_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyqueryparticipatenumberResponse `json:"result,omitempty" xml:"result,omitempty"`
}
