package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligeniealbumsgetAPIResponse 专辑详情 API返回值
// alibaba.ailabs.aligenie.albums.get
//
// 给予厂商查询专辑下的音频详情
type AlibabaailabsaligeniealbumsgetAPIResponse struct {
	model.CommonResponse
	AlibabaailabsaligeniealbumsgetAPIResponseModel
}

// AlibabaailabsaligeniealbumsgetAPIResponseModel is 专辑详情 成功返回结果
type AlibabaailabsaligeniealbumsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_albums_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
