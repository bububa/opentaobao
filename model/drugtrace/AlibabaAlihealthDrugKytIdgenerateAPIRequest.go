package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytIdgenerateAPIRequest
终端(医疗机构|零售药店)ID生成接口 API请求
alibaba.alihealth.drug.kyt.idgenerate

终端(医疗机构|零售药店)ID生成接口 */
type AlibabaAlihealthDrugKytIdgenerateAPIRequest struct {
	model.Params
	// 行政区（省市区）
	_regionCode string
	// 零售药店、医疗机构名称
	_terminalName string
}

// New
