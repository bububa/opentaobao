package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpbidwordsuggestdefaultlistAPIResponse 建议默认关键词 API返回值
// taobao.universalbp.bidword.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词
type TaobaouniversalbpbidwordsuggestdefaultlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpbidwordsuggestdefaultlistAPIResponseModel
}

// TaobaouniversalbpbidwordsuggestdefaultlistAPIResponseModel is 建议默认关键词 成功返回结果
type TaobaouniversalbpbidwordsuggestdefaultlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_bidword_suggestdefaultlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpbidwordsuggestdefaultlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
