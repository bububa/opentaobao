package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpwordpackagesuggestdefaultlistAPIResponse 建议默认关键词包 API返回值
// taobao.universalbp.wordpackage.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词包
type TaobaouniversalbpwordpackagesuggestdefaultlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpwordpackagesuggestdefaultlistAPIResponseModel
}

// TaobaouniversalbpwordpackagesuggestdefaultlistAPIResponseModel is 建议默认关键词包 成功返回结果
type TaobaouniversalbpwordpackagesuggestdefaultlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_wordpackage_suggestdefaultlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpwordpackagesuggestdefaultlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
