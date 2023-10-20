package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqafscheck 反欺诈二次验证接口
// alibaba.security.jaq.afs.check
//
// 反欺诈二次验证接口
func Alibabasecurityjaqafscheck(clt *core.SDKClient, req *security.AlibabasecurityjaqafscheckAPIRequest, session string) (*security.AlibabasecurityjaqafscheckAPIResponse, error) {
	var resp security.AlibabasecurityjaqafscheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
