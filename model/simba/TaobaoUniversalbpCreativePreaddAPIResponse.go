package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativePreaddAPIResponse 创建单品创意前置信息 API返回值
// taobao.universalbp.creative.preadd
//
// 用于关键词场景创建单品创意前使用
type TaobaoUniversalbpCreativePreaddAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCreativePreaddAPIResponseModel
}

// TaobaoUniversalbpCreativePreaddAPIResponseModel is 创建单品创意前置信息 成功返回结果
type TaobaoUniversalbpCreativePreaddAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_creative_preadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCreativePreaddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
