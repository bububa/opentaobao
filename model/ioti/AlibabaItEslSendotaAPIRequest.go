package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItEslSendotaAPIRequest
电子价签ota接口 API请求
alibaba.it.esl.sendota

厂测接口，电子价签ota接口 */
type AlibabaItEslSendotaAPIRequest struct {
	model.Params
	// mac
	_macAp string
	// base64的ota包
	_otaDataBase64String string
}

// New
