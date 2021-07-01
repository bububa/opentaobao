package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest
药厂传输数据 API请求
alibaba.alihealth.drugcode.drugfactory.transferdata

药厂传输数据 */
type AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest struct {
	model.Params
	// 时间戳(毫秒级别)
	_timestampYl int64
	// 签名值
	_signValue string
	// 密文
	_cipherText string
	// 企业Id
	_refEntId string
}

// New
