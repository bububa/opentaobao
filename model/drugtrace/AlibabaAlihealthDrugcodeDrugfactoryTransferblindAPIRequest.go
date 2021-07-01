package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest
传输盲底文件 API请求
alibaba.alihealth.drugcode.drugfactory.transferblind

临床用药试验-传输盲底文件 */
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 签名值
	_signValue string
	// 密文
	_cipherText string
	// 文件名称
	_fileName string
}

// New
