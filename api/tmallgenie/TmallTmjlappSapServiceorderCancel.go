package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Tmalltmjlappsapserviceordercancel 取消售后服务单
// tmall.tmjlapp.sap.serviceorder.cancel
//
// SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
func Tmalltmjlappsapserviceordercancel(clt *core.SDKClient, req *tmallgenie.TmalltmjlappsapserviceordercancelAPIRequest, session string) (*tmallgenie.TmalltmjlappsapserviceordercancelAPIResponse, error) {
	var resp tmallgenie.TmalltmjlappsapserviceordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
