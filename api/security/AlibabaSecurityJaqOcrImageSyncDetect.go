package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqocrimagesyncdetect 聚安全图文识别同步检测接口
// alibaba.security.jaq.ocr.image.sync.detect
//
// 图像字符识别同步检测接口
func Alibabasecurityjaqocrimagesyncdetect(clt *core.SDKClient, req *security.AlibabasecurityjaqocrimagesyncdetectAPIRequest, session string) (*security.AlibabasecurityjaqocrimagesyncdetectAPIResponse, error) {
	var resp security.AlibabasecurityjaqocrimagesyncdetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
