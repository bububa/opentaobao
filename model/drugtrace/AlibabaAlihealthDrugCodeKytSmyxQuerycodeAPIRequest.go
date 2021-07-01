package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest
扫码营销码查询 API请求
alibaba.alihealth.drug.code.kyt.smyx.querycode

此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
核查平台优先过滤非8开头的，长度非20位数字的码信息。 */
type AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 码列表
	_codes []string
}

// New
