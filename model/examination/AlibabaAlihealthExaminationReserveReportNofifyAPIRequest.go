package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveReportNofifyAPIRequest
服务商主动通知体检报告 API请求
alibaba.alihealth.examination.reserve.report.nofify

服务商主动回传用户的体检报告数据 */
type AlibabaAlihealthExaminationReserveReportNofifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 服务商到检编号
	_checkNo string
	// 健康预约凭证
	_reserveNumber string
	// 报告通知类型，传1即可
	_type string
	// pdf文件的二进制base64编码字符串
	_content string
}

// New
