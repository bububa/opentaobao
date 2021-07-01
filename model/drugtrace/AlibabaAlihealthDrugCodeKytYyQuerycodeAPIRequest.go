package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest
医院根据码查询码信息 API请求
alibaba.alihealth.drug.code.kyt.yy.querycode

服务描述
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
核查平台优先过滤非8开头的，长度非20位数字的码信息。 */
type AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest struct {
	model.Params
	// 企业唯一标识（或appkey）
	_refEntId string
	// 码列表
	_codes []string
}

// New
