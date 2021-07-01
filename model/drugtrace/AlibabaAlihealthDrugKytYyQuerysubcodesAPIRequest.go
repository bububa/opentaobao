package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest
查询一个码的所有子码 API请求
alibaba.alihealth.drug.kyt.yy.querysubcodes

单码的了码查询 */
type AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 码
	_codes []string
}

// New
