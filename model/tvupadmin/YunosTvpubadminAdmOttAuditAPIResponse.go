package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminAdmOttAuditAPIResponse 优酷OTT广告素材审核 API返回值
// yunos.tvpubadmin.adm.ott.audit
//
// 审核优酷OTT端广告素材
type YunosTvpubadminAdmOttAuditAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminAdmOttAuditAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminAdmOttAuditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminAdmOttAuditAPIResponseModel).Reset()
}

// YunosTvpubadminAdmOttAuditAPIResponseModel is 优酷OTT广告素材审核 成功返回结果
type YunosTvpubadminAdmOttAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_adm_ott_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的操作结果
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminAdmOttAuditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminAdmOttAuditAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminAdmOttAuditAPIResponse)
	},
}

// GetYunosTvpubadminAdmOttAuditAPIResponse 从 sync.Pool 获取 YunosTvpubadminAdmOttAuditAPIResponse
func GetYunosTvpubadminAdmOttAuditAPIResponse() *YunosTvpubadminAdmOttAuditAPIResponse {
	return poolYunosTvpubadminAdmOttAuditAPIResponse.Get().(*YunosTvpubadminAdmOttAuditAPIResponse)
}

// ReleaseYunosTvpubadminAdmOttAuditAPIResponse 将 YunosTvpubadminAdmOttAuditAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminAdmOttAuditAPIResponse(v *YunosTvpubadminAdmOttAuditAPIResponse) {
	v.Reset()
	poolYunosTvpubadminAdmOttAuditAPIResponse.Put(v)
}
