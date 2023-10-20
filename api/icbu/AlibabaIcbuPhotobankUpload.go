package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuphotobankupload 图片银行图片上传开放接口
// alibaba.icbu.photobank.upload
//
// 图片银行图片上传开放接口
func Alibabaicbuphotobankupload(clt *core.SDKClient, req *icbu.AlibabaicbuphotobankuploadAPIRequest, session string) (*icbu.AlibabaicbuphotobankuploadAPIResponse, error) {
	var resp icbu.AlibabaicbuphotobankuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
