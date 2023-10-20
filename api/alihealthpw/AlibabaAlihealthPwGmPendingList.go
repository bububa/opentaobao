package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// Alibabaalihealthpwgmpendinglist 同情用药待审核工单查询接口
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
func Alibabaalihealthpwgmpendinglist(clt *core.SDKClient, req *alihealthpw.AlibabaalihealthpwgmpendinglistAPIRequest, session string) (*alihealthpw.AlibabaalihealthpwgmpendinglistAPIResponse, error) {
	var resp alihealthpw.AlibabaalihealthpwgmpendinglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
