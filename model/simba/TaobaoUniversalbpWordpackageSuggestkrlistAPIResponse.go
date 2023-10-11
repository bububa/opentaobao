package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpWordpackageSuggestkrlistAPIResponse 关键词包建议 API返回值
// taobao.universalbp.wordpackage.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词包
type TaobaoUniversalbpWordpackageSuggestkrlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpWordpackageSuggestkrlistAPIResponseModel
}

// TaobaoUniversalbpWordpackageSuggestkrlistAPIResponseModel is 关键词包建议 成功返回结果
type TaobaoUniversalbpWordpackageSuggestkrlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_wordpackage_suggestkrlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpWordpackageSuggestkrlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
