package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpLabelFindconfiglistAPIResponse 查询可用标签id信息 API返回值
// taobao.universalbp.label.findconfiglist
//
// 入参账号信息，出参可用标签id，用于下游接口入参
type TaobaoUniversalbpLabelFindconfiglistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpLabelFindconfiglistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpLabelFindconfiglistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpLabelFindconfiglistAPIResponseModel).Reset()
}

// TaobaoUniversalbpLabelFindconfiglistAPIResponseModel is 查询可用标签id信息 成功返回结果
type TaobaoUniversalbpLabelFindconfiglistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_label_findconfiglist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpLabelFindconfiglistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpLabelFindconfiglistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpLabelFindconfiglistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpLabelFindconfiglistAPIResponse)
	},
}

// GetTaobaoUniversalbpLabelFindconfiglistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpLabelFindconfiglistAPIResponse
func GetTaobaoUniversalbpLabelFindconfiglistAPIResponse() *TaobaoUniversalbpLabelFindconfiglistAPIResponse {
	return poolTaobaoUniversalbpLabelFindconfiglistAPIResponse.Get().(*TaobaoUniversalbpLabelFindconfiglistAPIResponse)
}

// ReleaseTaobaoUniversalbpLabelFindconfiglistAPIResponse 将 TaobaoUniversalbpLabelFindconfiglistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpLabelFindconfiglistAPIResponse(v *TaobaoUniversalbpLabelFindconfiglistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpLabelFindconfiglistAPIResponse.Put(v)
}
