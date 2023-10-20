package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenpointoperateAPIResponse 积分操作接口 API返回值
// alibaba.alsc.crm.open.point.operate
//
// 同步积分接口
type AlibabaalsccrmopenpointoperateAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmopenpointoperateAPIResponseModel
}

// AlibabaalsccrmopenpointoperateAPIResponseModel is 积分操作接口 成功返回结果
type AlibabaalsccrmopenpointoperateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_point_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
