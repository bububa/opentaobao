package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationAgreementListAPIRequest
isv协议获取 API请求
alibaba.alihealth.examination.agreement.list

isv协议获取 */
type AlibabaAlihealthExaminationAgreementListAPIRequest struct {
	model.Params
	// isv传递过来的门店code
	_storeCode string
}

// New
