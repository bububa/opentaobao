package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveStateAPIRequest
体检机构对接_体检状态查询 API请求
alibaba.alihealth.examination.reserve.state

体检机构对接_体检状态查询 */
type AlibabaAlihealthExaminationReserveStateAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
}

// New
