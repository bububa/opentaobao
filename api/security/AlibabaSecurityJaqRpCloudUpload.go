package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpcloudupload 实人认证云上传接口
// alibaba.security.jaq.rp.cloud.upload
//
// 聚安全实人认证上传认证信息
func Alibabasecurityjaqrpcloudupload(clt *core.SDKClient, req *security.AlibabasecurityjaqrpclouduploadAPIRequest, session string) (*security.AlibabasecurityjaqrpclouduploadAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpclouduploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
