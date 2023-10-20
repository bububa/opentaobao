package customizemarket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabataobaoindustrypetprofilequeryAPIResponse 用户宠物列表查询 API返回值
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
type AlibabataobaoindustrypetprofilequeryAPIResponse struct {
	model.CommonResponse
	AlibabataobaoindustrypetprofilequeryAPIResponseModel
}

// AlibabataobaoindustrypetprofilequeryAPIResponseModel is 用户宠物列表查询 成功返回结果
type AlibabataobaoindustrypetprofilequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_taobao_industry_pet_profile_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数错误
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 501
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 数据
	Object *BasePageBean `json:"object,omitempty" xml:"object,omitempty"`
	// 是否成功
	ResultStatus bool `json:"result_status,omitempty" xml:"result_status,omitempty"`
}
