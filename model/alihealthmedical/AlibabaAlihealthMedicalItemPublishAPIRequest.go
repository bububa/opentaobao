package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalItemPublishAPIRequest
三方入驻-开通服务 API请求
alibaba.alihealth.medical.item.publish

三方入驻-开通服务 */
type AlibabaAlihealthMedicalItemPublishAPIRequest struct {
	model.Params
	// 请求
	_request1 *ItemPublishRequest
}

// New
