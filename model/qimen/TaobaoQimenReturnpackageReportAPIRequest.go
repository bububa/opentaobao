package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenReturnpackageReportAPIRequest
退货包裹状态通知接口 API请求
taobao.qimen.returnpackage.report

退货包裹状态通知接口 */
type TaobaoQimenReturnpackageReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenReturnpackageReportRequest
}

// New
