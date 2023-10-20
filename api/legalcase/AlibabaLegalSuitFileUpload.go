package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalsuitfileupload 诉讼文件上传接口
// alibaba.legal.suit.file.upload
//
// 上传文件接口
func Alibabalegalsuitfileupload(clt *core.SDKClient, req *legalcase.AlibabalegalsuitfileuploadAPIRequest, session string) (*legalcase.AlibabalegalsuitfileuploadAPIResponse, error) {
	var resp legalcase.AlibabalegalsuitfileuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
