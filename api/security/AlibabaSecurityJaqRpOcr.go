package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpocr 聚安全实人认证证件OCR识别功能接口
// alibaba.security.jaq.rp.ocr
//
// 聚安全实人认证证件OCR识别功能接口
func Alibabasecurityjaqrpocr(clt *core.SDKClient, req *security.AlibabasecurityjaqrpocrAPIRequest, session string) (*security.AlibabasecurityjaqrpocrAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpocrAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
