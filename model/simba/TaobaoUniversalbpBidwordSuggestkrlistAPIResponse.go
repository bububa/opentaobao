package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpbidwordsuggestkrlistAPIResponse 关键词建议 API返回值
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
type TaobaouniversalbpbidwordsuggestkrlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpbidwordsuggestkrlistAPIResponseModel
}

// TaobaouniversalbpbidwordsuggestkrlistAPIResponseModel is 关键词建议 成功返回结果
type TaobaouniversalbpbidwordsuggestkrlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_bidword_suggestkrlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpbidwordsuggestkrlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
