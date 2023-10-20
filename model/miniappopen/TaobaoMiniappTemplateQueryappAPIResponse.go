package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateQueryappAPIResponse 查询实例化应用版本 API返回值
// taobao.miniapp.template.queryapp
//
// 根据模板id和商家信息，查询实例化小程序版本查询
type TaobaoMiniappTemplateQueryappAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappTemplateQueryappAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateQueryappAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappTemplateQueryappAPIResponseModel).Reset()
}

// TaobaoMiniappTemplateQueryappAPIResponseModel is 查询实例化应用版本 成功返回结果
type TaobaoMiniappTemplateQueryappAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_template_queryapp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 所有版本信息
	AllVersionInfos []MiniappInstanceAppAllVersionsDto `json:"all_version_infos,omitempty" xml:"all_version_infos>miniapp_instance_app_all_versions_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateQueryappAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AllVersionInfos = m.AllVersionInfos[:0]
}

var poolTaobaoMiniappTemplateQueryappAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappTemplateQueryappAPIResponse)
	},
}

// GetTaobaoMiniappTemplateQueryappAPIResponse 从 sync.Pool 获取 TaobaoMiniappTemplateQueryappAPIResponse
func GetTaobaoMiniappTemplateQueryappAPIResponse() *TaobaoMiniappTemplateQueryappAPIResponse {
	return poolTaobaoMiniappTemplateQueryappAPIResponse.Get().(*TaobaoMiniappTemplateQueryappAPIResponse)
}

// ReleaseTaobaoMiniappTemplateQueryappAPIResponse 将 TaobaoMiniappTemplateQueryappAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappTemplateQueryappAPIResponse(v *TaobaoMiniappTemplateQueryappAPIResponse) {
	v.Reset()
	poolTaobaoMiniappTemplateQueryappAPIResponse.Put(v)
}
