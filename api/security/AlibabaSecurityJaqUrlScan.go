package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqurlscan 恶意网址检测接口
// alibaba.security.jaq.url.scan
//
// url扫描接口
func Alibabasecurityjaqurlscan(clt *core.SDKClient, req *security.AlibabasecurityjaqurlscanAPIRequest, session string) (*security.AlibabasecurityjaqurlscanAPIResponse, error) {
	var resp security.AlibabasecurityjaqurlscanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
