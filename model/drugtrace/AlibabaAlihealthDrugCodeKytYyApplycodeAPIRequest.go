package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest
医院药品子码申请接口 API请求
alibaba.alihealth.drug.code.kyt.yy.applycode

根据父码及所属企业ID生成子码信息 */
type AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest struct {
	model.Params
	// 企业ID（ref_ent_id)
	_refEntId string
	// 父码
	_code string
	// 申请数量
	_amount int64
}

// New
