package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuSpBillreordAddAPIRequest
内购服务确认单明细上传接口 API请求
taobao.fuwu.sp.billreord.add

isv能通过该接口上传确认单明细数据 */
type TaobaoFuwuSpBillreordAddAPIRequest struct {
	model.Params
	// 确认单的账单明细
	_paramBillRecordDTO *BillRecordDto
}

// New
