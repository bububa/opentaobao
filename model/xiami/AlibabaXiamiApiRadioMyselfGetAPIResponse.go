package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaxiamiapiradiomyselfgetAPIResponse 我的电台 API返回值
// alibaba.xiami.api.radio.myself.get
//
// 我的电台
type AlibabaxiamiapiradiomyselfgetAPIResponse struct {
	model.CommonResponse
	AlibabaxiamiapiradiomyselfgetAPIResponseModel
}

// AlibabaxiamiapiradiomyselfgetAPIResponseModel is 我的电台 成功返回结果
type AlibabaxiamiapiradiomyselfgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_radio_myself_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲列表
	Data []Song `json:"data,omitempty" xml:"data>song,omitempty"`
}
