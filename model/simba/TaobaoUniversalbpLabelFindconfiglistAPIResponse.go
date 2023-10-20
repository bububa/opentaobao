package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbplabelfindconfiglistAPIResponse 查询可用标签id信息 API返回值
// taobao.universalbp.label.findconfiglist
//
// 入参账号信息，出参可用标签id，用于下游接口入参
type TaobaouniversalbplabelfindconfiglistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbplabelfindconfiglistAPIResponseModel
}

// TaobaouniversalbplabelfindconfiglistAPIResponseModel is 查询可用标签id信息 成功返回结果
type TaobaouniversalbplabelfindconfiglistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_label_findconfiglist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbplabelfindconfiglistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
