package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwApplynodeUpdatename 回调变更患者姓名
// alibaba.alihealth.pw.applynode.updatename
//
// 回调变更患者姓名
func AlibabaAlihealthPwApplynodeUpdatename(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwApplynodeUpdatenameAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwApplynodeUpdatenameAPIResponse, error) {
	var resp alihealthpw.AlibabaAlihealthPwApplynodeUpdatenameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
