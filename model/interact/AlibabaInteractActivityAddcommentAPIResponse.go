package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityAddcommentAPIResponse 微淘评论接口 API返回值
// alibaba.interact.activity.addcomment
//
// 发表评论，并返回楼层
type AlibabaInteractActivityAddcommentAPIResponse struct {
	model.CommonResponse
	AlibabaInteractActivityAddcommentAPIResponseModel
}

// AlibabaInteractActivityAddcommentAPIResponseModel is 微淘评论接口 成功返回结果
type AlibabaInteractActivityAddcommentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_activity_addcomment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 评论的楼层数
	Floor int64 `json:"floor,omitempty" xml:"floor,omitempty"`
	// 返回成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
