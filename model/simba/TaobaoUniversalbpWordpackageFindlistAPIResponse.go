package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpwordpackagefindlistAPIResponse 词包列表查询 API返回值
// taobao.universalbp.wordpackage.findlist
//
// 根据计划+单元id，查绑定的词包列表
type TaobaouniversalbpwordpackagefindlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpwordpackagefindlistAPIResponseModel
}

// TaobaouniversalbpwordpackagefindlistAPIResponseModel is 词包列表查询 成功返回结果
type TaobaouniversalbpwordpackagefindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_wordpackage_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpwordpackagefindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
