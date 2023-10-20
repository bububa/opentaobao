package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqappofficialverify 聚安全验证官方应用接口
// alibaba.security.jaq.app.official.verify
//
// 接入用户来查询应用是否为官方应用
func Alibabasecurityjaqappofficialverify(clt *core.SDKClient, req *security.AlibabasecurityjaqappofficialverifyAPIRequest, session string) (*security.AlibabasecurityjaqappofficialverifyAPIResponse, error) {
	var resp security.AlibabasecurityjaqappofficialverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
