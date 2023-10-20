package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse 商家基础经营设置信息同步 API返回值
// taobao.koubei.saas.base.operation.config.sync
//
// ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS
type TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponseModel).Reset()
}

// TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponseModel is 商家基础经营设置信息同步 成功返回结果
type TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_saas_base_operation_config_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 异常码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.BizSuccess = false
}

var poolTaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse)
	},
}

// GetTaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse 从 sync.Pool 获取 TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse
func GetTaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse() *TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse {
	return poolTaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse.Get().(*TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse)
}

// ReleaseTaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse 将 TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse(v *TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse.Put(v)
}
