package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// Alibabaalihealthpwapplynodeupdatename 回调变更患者姓名
// alibaba.alihealth.pw.applynode.updatename
//
// 回调变更患者姓名
func Alibabaalihealthpwapplynodeupdatename(clt *core.SDKClient, req *alihealthpw.AlibabaalihealthpwapplynodeupdatenameAPIRequest, session string) (*alihealthpw.AlibabaalihealthpwapplynodeupdatenameAPIResponse, error) {
	var resp alihealthpw.AlibabaalihealthpwapplynodeupdatenameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
