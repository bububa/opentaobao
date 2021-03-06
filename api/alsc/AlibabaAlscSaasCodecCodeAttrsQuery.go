package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscSaasCodecCodeAttrsQuery 码业务属性查询
// alibaba.alsc.saas.codec.code.attrs.query
//
// 码业务属性查询
func AlibabaAlscSaasCodecCodeAttrsQuery(clt *core.SDKClient, req *alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest, session string) (*alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse, error) {
	var resp alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
