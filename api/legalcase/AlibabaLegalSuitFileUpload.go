package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

/* AlibabaLegalSuitFileUpload
诉讼文件上传接口
alibaba.legal.suit.file.upload

上传文件接口 */
func AlibabaLegalSuitFileUpload(clt *core.SDKClient, req *legalcase.AlibabaLegalSuitFileUploadAPIRequest, session string) (*legalcase.AlibabaLegalSuitFileUploadAPIResponse, error) {
	var resp legalcase.AlibabaLegalSuitFileUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
