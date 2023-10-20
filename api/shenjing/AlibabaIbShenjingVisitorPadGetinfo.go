package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// Alibabaibshenjingvisitorpadgetinfo 获取OSS上传参数
// alibaba.ib.shenjing.visitor.pad.getinfo
//
// PAD 端获取OSS上传参数，向OSS服务器上传图片。
func Alibabaibshenjingvisitorpadgetinfo(clt *core.SDKClient, req *shenjing.AlibabaibshenjingvisitorpadgetinfoAPIRequest, session string) (*shenjing.AlibabaibshenjingvisitorpadgetinfoAPIResponse, error) {
	var resp shenjing.AlibabaibshenjingvisitorpadgetinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
