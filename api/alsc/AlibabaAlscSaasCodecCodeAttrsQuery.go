package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscsaascodeccodeattrsquery 码业务属性查询
// alibaba.alsc.saas.codec.code.attrs.query
//
// 码业务属性查询
func Alibabaalscsaascodeccodeattrsquery(clt *core.SDKClient, req *alsc.AlibabaalscsaascodeccodeattrsqueryAPIRequest, session string) (*alsc.AlibabaalscsaascodeccodeattrsqueryAPIResponse, error) {
	var resp alsc.AlibabaalscsaascodeccodeattrsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
