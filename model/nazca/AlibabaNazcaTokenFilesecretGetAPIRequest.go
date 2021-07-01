package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaTokenFilesecretGetAPIRequest
获取文件秘钥 API请求
alibaba.nazca.token.filesecret.get

获取文件秘钥 */
type AlibabaNazcaTokenFilesecretGetAPIRequest struct {
	model.Params
	// 客户在1688的唯一标识
	_platformUserId string
	// 合同编号
	_contractNum string
}

// New
