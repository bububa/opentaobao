package tvupadmin

import (
	"encoding/xml"

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

// YunosTvpubadminAdmOttAuditAPIResponseModel is 优酷OTT广告素材审核 成功返回结果
type YunosTvpubadminAdmOttAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_adm_ott_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的操作结果
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}
