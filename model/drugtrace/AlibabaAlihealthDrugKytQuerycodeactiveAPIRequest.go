package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest
查询码是否激活 API请求
alibaba.alihealth.drug.kyt.querycodeactive

查询码是否激活 */
type AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest struct {
	model.Params
	// 企业
	_refEntId string
	// 码
	_codes []string
}

// New
