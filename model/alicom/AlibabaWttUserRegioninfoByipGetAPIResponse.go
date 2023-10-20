package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWttUserRegioninfoByipGetAPIResponse 根据ip获取省市信息 API返回值
// alibaba.wtt.user.regioninfo.byip.get
//
// 通过ip获取省市信息
type AlibabaWttUserRegioninfoByipGetAPIResponse struct {
	model.CommonResponse
	AlibabaWttUserRegioninfoByipGetAPIResponseModel
}

// AlibabaWttUserRegioninfoByipGetAPIResponseModel is 根据ip获取省市信息 成功返回结果
type AlibabaWttUserRegioninfoByipGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wtt_user_regioninfo_byip_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地址信息
	Model *RegionInfo `json:"model,omitempty" xml:"model,omitempty"`
}
