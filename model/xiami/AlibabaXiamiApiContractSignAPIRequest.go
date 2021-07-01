package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiContractSignAPIRequest
提供签约链接 API请求
alibaba.xiami.api.contract.sign

提供签约链接。in：商家id；out：签约url */
type AlibabaXiamiApiContractSignAPIRequest struct {
	model.Params
}

// New
