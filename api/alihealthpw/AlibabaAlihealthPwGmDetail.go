package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwGmDetail 同情用药申请单详情接口
// alibaba.alihealth.pw.gm.detail
//
// 同情用药申请单详情接口，提供给合作方查询申请单详情
func AlibabaAlihealthPwGmDetail(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwGmDetailAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwGmDetailAPIResponse, error) {
	var resp alihealthpw.AlibabaAlihealthPwGmDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
