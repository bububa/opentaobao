package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqappofficialapply 聚安全官方应用申请
// alibaba.security.jaq.app.official.apply
//
// 官方应用申请接口
func Alibabasecurityjaqappofficialapply(clt *core.SDKClient, req *security.AlibabasecurityjaqappofficialapplyAPIRequest, session string) (*security.AlibabasecurityjaqappofficialapplyAPIResponse, error) {
	var resp security.AlibabasecurityjaqappofficialapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
