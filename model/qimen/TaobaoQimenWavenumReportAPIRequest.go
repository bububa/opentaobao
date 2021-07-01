package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenWavenumReportAPIRequest
发货单波次通知接口 API请求
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求 */
type TaobaoQimenWavenumReportAPIRequest struct {
	model.Params
	//
	_request *WaveNumReportRequest
}

// New
