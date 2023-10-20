package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpcloudrphit 实人认证云服务日志打点
// alibaba.security.jaq.rp.cloud.rphit
//
// 聚安全实人认证日志打点接口
func Alibabasecurityjaqrpcloudrphit(clt *core.SDKClient, req *security.AlibabasecurityjaqrpcloudrphitAPIRequest, session string) (*security.AlibabasecurityjaqrpcloudrphitAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpcloudrphitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
