package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtophotwordsgetAPIResponse 获取热词 API返回值
// taobao.ailab.aicloud.top.hotwords.get
//
// 获取ASR热词
type TaobaoailabaicloudtophotwordsgetAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtophotwordsgetAPIResponseModel
}

// TaobaoailabaicloudtophotwordsgetAPIResponseModel is 获取热词 成功返回结果
type TaobaoailabaicloudtophotwordsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_hotwords_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// baseresult
	Baseresult *BaseResult `json:"baseresult,omitempty" xml:"baseresult,omitempty"`
}
