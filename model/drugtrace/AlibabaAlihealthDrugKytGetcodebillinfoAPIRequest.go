package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest
根据码获取基本和单据信息 API请求
alibaba.alihealth.drug.kyt.getcodebillinfo

根据码信息获取基本信息和单据信息 */
type AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 码
	_code string
}

// New
