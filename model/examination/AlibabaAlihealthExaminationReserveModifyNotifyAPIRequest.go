package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest
通知改期结果 API请求
alibaba.alihealth.examination.reserve.modify.notify

体检状态为改期中，服务上通知健康是否改期成功 */
type AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest struct {
	model.Params
	// 旧的预约日期
	_oldReserveDate string
	// 套餐编码
	_packageCode string
	// 健康预约凭证
	_reserveNumber string
	// 新的预约日期
	_newReserveDate string
	// 服务商预约凭证
	_uniqReserveCode string
	// 商品编码
	_goodsCode string
	// 门店编码
	_storeCode string
	// true:同意修改；false:拒绝修改
	_pass bool
	// 拒绝修改的时候需要传递拒绝原因
	_reason string
	// 新的预约时间段开始时间
	_newReserveTimeStart string
	// 新的预约时间段结束时间
	_newReserveTimeEnd string
}

// New
