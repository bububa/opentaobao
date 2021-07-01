package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeKytYqQuerycodeAPIRequest
查询追溯码对应的药品信息（疫情） API请求
alibaba.alihealth.drug.code.kyt.yq.querycode

通过追溯码码得到 药品名称、包装规格、剂型、剂型规格”、有效期至等信息。 */
type AlibabaAlihealthDrugCodeKytYqQuerycodeAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// 调用零售药店名称
	_terminalName string
	// 门店所属的行政区域ID
	_bureauId string
}

// New
