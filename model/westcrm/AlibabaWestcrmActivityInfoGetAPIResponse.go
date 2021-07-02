package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmActivityInfoGetAPIResponse 获取活动详情 API返回值
// alibaba.westcrm.activity.info.get
//
// 根据id查询活动详情
type AlibabaWestcrmActivityInfoGetAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmActivityInfoGetAPIResponseModel
}

// AlibabaWestcrmActivityInfoGetAPIResponseModel is 获取活动详情 成功返回结果
type AlibabaWestcrmActivityInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_activity_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
