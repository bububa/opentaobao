package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpbidwordfindlistAPIResponse 词列表查询 API返回值
// taobao.universalbp.bidword.findlist
//
// 根据计划+单元id，查绑定的词列表
type TaobaouniversalbpbidwordfindlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpbidwordfindlistAPIResponseModel
}

// TaobaouniversalbpbidwordfindlistAPIResponseModel is 词列表查询 成功返回结果
type TaobaouniversalbpbidwordfindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_bidword_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpbidwordfindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
