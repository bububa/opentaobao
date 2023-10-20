package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqpornimagesyncdetect 聚安全智能鉴黄同步检测接口
// alibaba.security.jaq.porn.image.sync.detect
//
// 同步黄图图像检测接口
func Alibabasecurityjaqpornimagesyncdetect(clt *core.SDKClient, req *security.AlibabasecurityjaqpornimagesyncdetectAPIRequest, session string) (*security.AlibabasecurityjaqpornimagesyncdetectAPIResponse, error) {
	var resp security.AlibabasecurityjaqpornimagesyncdetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
