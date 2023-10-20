package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse 横向管理创意分页查询 API返回值
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
type TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel
}

// TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel is 横向管理创意分页查询 成功返回结果
type TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_creative_horizontal_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCreativeHorizontalFindpageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
