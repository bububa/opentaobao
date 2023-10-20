package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpocrcheck ocr同时实名校验
// alibaba.security.jaq.rp.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
func Alibabasecurityjaqrpocrcheck(clt *core.SDKClient, req *security.AlibabasecurityjaqrpocrcheckAPIRequest, session string) (*security.AlibabasecurityjaqrpocrcheckAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpocrcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
