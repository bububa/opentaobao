package traderate

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmalltraderateitemtagsgetAPIResponse 通过商品ID获取标签列表 API返回值
// tmall.traderate.itemtags.get
//
// 通过商品ID获取标签详细信息
type TmalltraderateitemtagsgetAPIResponse struct {
	model.CommonResponse
	TmalltraderateitemtagsgetAPIResponseModel
}

// TmalltraderateitemtagsgetAPIResponseModel is 通过商品ID获取标签列表 成功返回结果
type TmalltraderateitemtagsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traderate_itemtags_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标签列表
	Tags []TmallRateTagDetail `json:"tags,omitempty" xml:"tags>tmall_rate_tag_detail,omitempty"`
}
