package alitrippoi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformContentRawAddAPIResponse 穷游内容写入接口 API返回值
// alitrip.platform.content.raw.add
//
// 穷游内容写入飞猪接口
type AlitripPlatformContentRawAddAPIResponse struct {
	model.CommonResponse
	AlitripPlatformContentRawAddAPIResponseModel
}

// AlitripPlatformContentRawAddAPIResponseModel is 穷游内容写入接口 成功返回结果
type AlitripPlatformContentRawAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_content_raw_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
