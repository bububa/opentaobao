package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// Alibabaicbuannexupload 上传附件获取附件files_str
// alibaba.icbu.annex.upload
//
// 上传附件获取附件files_str
func Alibabaicbuannexupload(clt *core.SDKClient, req *icburfq.AlibabaicbuannexuploadAPIRequest, session string) (*icburfq.AlibabaicbuannexuploadAPIResponse, error) {
	var resp icburfq.AlibabaicbuannexuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
