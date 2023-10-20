package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// Alibabaibshenjingvisitorpaduploadface 访客PAD上传人脸
// alibaba.ib.shenjing.visitor.pad.uploadface
//
// 访客PAD端上传人脸。
func Alibabaibshenjingvisitorpaduploadface(clt *core.SDKClient, req *shenjing.AlibabaibshenjingvisitorpaduploadfaceAPIRequest, session string) (*shenjing.AlibabaibshenjingvisitorpaduploadfaceAPIResponse, error) {
	var resp shenjing.AlibabaibshenjingvisitorpaduploadfaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
