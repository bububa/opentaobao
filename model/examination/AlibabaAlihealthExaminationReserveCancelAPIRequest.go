package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveCancelAPIRequest
体检机构对接_预约取消 API请求
alibaba.alihealth.examination.reserve.cancel

体检机构对接_体检取消 */
type AlibabaAlihealthExaminationReserveCancelAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 预约时间
	_reserveDate string
	// 体检套餐编码
	_packageCode string
	// 店铺ID
	_storeId string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
}

// New
