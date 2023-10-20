package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSearchtagtemplateGetAPIResponse 获取搜索人群TOP用户可添加人群信息 API返回值
// taobao.simba.searchtagtemplate.get
//
// 获取搜索人群用户可添加人群信息
type TaobaoSimbaSearchtagtemplateGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSearchtagtemplateGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSearchtagtemplateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSearchtagtemplateGetAPIResponseModel).Reset()
}

// TaobaoSimbaSearchtagtemplateGetAPIResponseModel is 获取搜索人群TOP用户可添加人群信息 成功返回结果
type TaobaoSimbaSearchtagtemplateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_searchtagtemplate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	TemplateList []TaobaoSimbaSearchtagtemplateGetResult `json:"template_list,omitempty" xml:"template_list>taobao_simba_searchtagtemplate_get_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSearchtagtemplateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TemplateList = m.TemplateList[:0]
}

var poolTaobaoSimbaSearchtagtemplateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSearchtagtemplateGetAPIResponse)
	},
}

// GetTaobaoSimbaSearchtagtemplateGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaSearchtagtemplateGetAPIResponse
func GetTaobaoSimbaSearchtagtemplateGetAPIResponse() *TaobaoSimbaSearchtagtemplateGetAPIResponse {
	return poolTaobaoSimbaSearchtagtemplateGetAPIResponse.Get().(*TaobaoSimbaSearchtagtemplateGetAPIResponse)
}

// ReleaseTaobaoSimbaSearchtagtemplateGetAPIResponse 将 TaobaoSimbaSearchtagtemplateGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSearchtagtemplateGetAPIResponse(v *TaobaoSimbaSearchtagtemplateGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSearchtagtemplateGetAPIResponse.Put(v)
}
