package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaisassetstagabortAPIResponse 基础设施资产标签废弃 API返回值
// alibaba.ais.assets.tag.abort
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃
type AlibabaaisassetstagabortAPIResponse struct {
	model.CommonResponse
	AlibabaaisassetstagabortAPIResponseModel
}

// AlibabaaisassetstagabortAPIResponseModel is 基础设施资产标签废弃 成功返回结果
type AlibabaaisassetstagabortAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ais_assets_tag_abort_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *BaseRep `json:"result,omitempty" xml:"result,omitempty"`
}
