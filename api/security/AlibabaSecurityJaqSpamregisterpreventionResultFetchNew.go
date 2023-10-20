package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqspamregisterpreventionresultfetchnew 获取虚假注册保护结果
// alibaba.security.jaq.spamregisterprevention.result.fetch.new
//
// 获取虚假注册保护结果
func Alibabasecurityjaqspamregisterpreventionresultfetchnew(clt *core.SDKClient, req *security.AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest, session string) (*security.AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIResponse, error) {
	var resp security.AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
