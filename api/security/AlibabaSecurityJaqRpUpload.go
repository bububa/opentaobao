package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpupload 聚安全实人认证上传认证信息
// alibaba.security.jaq.rp.upload
//
// 聚安全实人认证上传认证信息
func Alibabasecurityjaqrpupload(clt *core.SDKClient, req *security.AlibabasecurityjaqrpuploadAPIRequest, session string) (*security.AlibabasecurityjaqrpuploadAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
