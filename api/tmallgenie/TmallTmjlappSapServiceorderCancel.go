package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TmallTmjlappSapServiceorderCancel 取消售后服务单
// tmall.tmjlapp.sap.serviceorder.cancel
//
// SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
func TmallTmjlappSapServiceorderCancel(clt *core.SDKClient, req *tmallgenie.TmallTmjlappSapServiceorderCancelAPIRequest, session string) (*tmallgenie.TmallTmjlappSapServiceorderCancelAPIResponse, error) {
	var resp tmallgenie.TmallTmjlappSapServiceorderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
