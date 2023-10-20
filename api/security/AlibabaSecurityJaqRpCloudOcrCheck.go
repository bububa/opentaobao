package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpcloudocrcheck ocr同时实名校验
// alibaba.security.jaq.rp.cloud.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
func Alibabasecurityjaqrpcloudocrcheck(clt *core.SDKClient, req *security.AlibabasecurityjaqrpcloudocrcheckAPIRequest, session string) (*security.AlibabasecurityjaqrpcloudocrcheckAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpcloudocrcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
