package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpfetchmaterial 聚安全实人认证获取结果接口
// alibaba.security.jaq.rp.fetchmaterial
//
// 聚安全实人认证获取结果接口
func Alibabasecurityjaqrpfetchmaterial(clt *core.SDKClient, req *security.AlibabasecurityjaqrpfetchmaterialAPIRequest, session string) (*security.AlibabasecurityjaqrpfetchmaterialAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpfetchmaterialAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
