package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpcloudevent 事件上报
// alibaba.security.jaq.rp.cloud.event
//
// 事件上报接口
func Alibabasecurityjaqrpcloudevent(clt *core.SDKClient, req *security.AlibabasecurityjaqrpcloudeventAPIRequest, session string) (*security.AlibabasecurityjaqrpcloudeventAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpcloudeventAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
