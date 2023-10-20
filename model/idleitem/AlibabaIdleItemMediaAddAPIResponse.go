package idleitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleitemmediaaddAPIResponse 图片上传 API返回值
// alibaba.idle.item.media.add
//
// 上传图片
type AlibabaidleitemmediaaddAPIResponse struct {
	model.CommonResponse
	AlibabaidleitemmediaaddAPIResponseModel
}

// AlibabaidleitemmediaaddAPIResponseModel is 图片上传 成功返回结果
type AlibabaidleitemmediaaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_item_media_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *EasyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
