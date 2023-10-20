package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabavalueusergetidAPIResponse 获取用户userId API返回值
// alibaba.value.user.getid
//
// 获取用户userId
type AlibabavalueusergetidAPIResponse struct {
	model.CommonResponse
	AlibabavalueusergetidAPIResponseModel
}

// AlibabavalueusergetidAPIResponseModel is 获取用户userId 成功返回结果
type AlibabavalueusergetidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_value_user_getid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	RespMsg string `json:"resp_msg,omitempty" xml:"resp_msg,omitempty"`
	// 错误码
	RespCode int64 `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 用户渠道开放信息
	Data *UserChannelOpenDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功状态
	SuccessStatus bool `json:"success_status,omitempty" xml:"success_status,omitempty"`
}
