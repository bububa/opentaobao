package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// Alibabaalihealthpwgmdetail 同情用药申请单详情接口
// alibaba.alihealth.pw.gm.detail
//
// 同情用药申请单详情接口，提供给合作方查询申请单详情
func Alibabaalihealthpwgmdetail(clt *core.SDKClient, req *alihealthpw.AlibabaalihealthpwgmdetailAPIRequest, session string) (*alihealthpw.AlibabaalihealthpwgmdetailAPIResponse, error) {
	var resp alihealthpw.AlibabaalihealthpwgmdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
