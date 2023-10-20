package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateOnlineappAPIResponse 上线实例化应用 API返回值
// taobao.miniapp.template.onlineapp
//
// 将指定的预览版本发布上线，预览版本号由构建实例化或更新实例化接口返回。
type TaobaoMiniappTemplateOnlineappAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappTemplateOnlineappAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateOnlineappAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappTemplateOnlineappAPIResponseModel).Reset()
}

// TaobaoMiniappTemplateOnlineappAPIResponseModel is 上线实例化应用 成功返回结果
type TaobaoMiniappTemplateOnlineappAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_template_onlineapp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分端上线结果
	OnlineResults []MiniappInstanceAppOnlineDto `json:"online_results,omitempty" xml:"online_results>miniapp_instance_app_online_dto,omitempty"`
	// 基本信息
	AppInfo *MiniAppEntityTemplateDto `json:"app_info,omitempty" xml:"app_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateOnlineappAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OnlineResults = m.OnlineResults[:0]
	m.AppInfo = nil
}

var poolTaobaoMiniappTemplateOnlineappAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappTemplateOnlineappAPIResponse)
	},
}

// GetTaobaoMiniappTemplateOnlineappAPIResponse 从 sync.Pool 获取 TaobaoMiniappTemplateOnlineappAPIResponse
func GetTaobaoMiniappTemplateOnlineappAPIResponse() *TaobaoMiniappTemplateOnlineappAPIResponse {
	return poolTaobaoMiniappTemplateOnlineappAPIResponse.Get().(*TaobaoMiniappTemplateOnlineappAPIResponse)
}

// ReleaseTaobaoMiniappTemplateOnlineappAPIResponse 将 TaobaoMiniappTemplateOnlineappAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappTemplateOnlineappAPIResponse(v *TaobaoMiniappTemplateOnlineappAPIResponse) {
	v.Reset()
	poolTaobaoMiniappTemplateOnlineappAPIResponse.Put(v)
}
