package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// Alibabaalihealthpwgmidslist 同情用药根据申请单列表查询申请单
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
func Alibabaalihealthpwgmidslist(clt *core.SDKClient, req *alihealthpw.AlibabaalihealthpwgmidslistAPIRequest, session string) (*alihealthpw.AlibabaalihealthpwgmidslistAPIResponse, error) {
	var resp alihealthpw.AlibabaalihealthpwgmidslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
