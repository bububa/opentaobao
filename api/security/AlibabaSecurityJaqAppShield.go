package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqappshield 应用加固接口
// alibaba.security.jaq.app.shield
//
// 提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
func Alibabasecurityjaqappshield(clt *core.SDKClient, req *security.AlibabasecurityjaqappshieldAPIRequest, session string) (*security.AlibabasecurityjaqappshieldAPIResponse, error) {
	var resp security.AlibabasecurityjaqappshieldAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
