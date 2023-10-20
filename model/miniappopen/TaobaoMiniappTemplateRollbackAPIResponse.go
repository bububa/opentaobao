package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateRollbackAPIResponse 回滚实例化应用 API返回值
// taobao.miniapp.template.rollback
//
// 将实例化小程序回滚到指定版本
type TaobaoMiniappTemplateRollbackAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappTemplateRollbackAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateRollbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappTemplateRollbackAPIResponseModel).Reset()
}

// TaobaoMiniappTemplateRollbackAPIResponseModel is 回滚实例化应用 成功返回结果
type TaobaoMiniappTemplateRollbackAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_template_rollback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分端回滚结果
	OnlineResults []MiniappInstanceAppOnlineDto `json:"online_results,omitempty" xml:"online_results>miniapp_instance_app_online_dto,omitempty"`
	// 基本信息
	AppInfo *MiniAppEntityTemplateDto `json:"app_info,omitempty" xml:"app_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateRollbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OnlineResults = m.OnlineResults[:0]
	m.AppInfo = nil
}

var poolTaobaoMiniappTemplateRollbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappTemplateRollbackAPIResponse)
	},
}

// GetTaobaoMiniappTemplateRollbackAPIResponse 从 sync.Pool 获取 TaobaoMiniappTemplateRollbackAPIResponse
func GetTaobaoMiniappTemplateRollbackAPIResponse() *TaobaoMiniappTemplateRollbackAPIResponse {
	return poolTaobaoMiniappTemplateRollbackAPIResponse.Get().(*TaobaoMiniappTemplateRollbackAPIResponse)
}

// ReleaseTaobaoMiniappTemplateRollbackAPIResponse 将 TaobaoMiniappTemplateRollbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappTemplateRollbackAPIResponse(v *TaobaoMiniappTemplateRollbackAPIResponse) {
	v.Reset()
	poolTaobaoMiniappTemplateRollbackAPIResponse.Put(v)
}
