package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest
获取加密公钥 API请求
alibaba.alihealth.drugcode.drugfactory.getencrptypk

获取服务端给药厂用来加密的公钥 */
type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest struct {
	model.Params
	// 企业Id
	_refEntId string
}

// New
