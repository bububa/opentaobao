package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpquery 聚安全实人认证查询认证结果
// alibaba.security.jaq.rp.query
//
// 聚安全实人认证查询认证结果
func Alibabasecurityjaqrpquery(clt *core.SDKClient, req *security.AlibabasecurityjaqrpqueryAPIRequest, session string) (*security.AlibabasecurityjaqrpqueryAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
