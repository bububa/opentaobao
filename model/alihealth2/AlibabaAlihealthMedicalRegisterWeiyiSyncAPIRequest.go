package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest
微医数据号源回传 API请求
alibaba.alihealth.medical.register.weiyi.sync

微医号源数据回传 */
type AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest struct {
	model.Params
	// 号源数据实体
	_serviceRequest *SourcesReturnVo
}

// New
