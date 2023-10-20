package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// Alibabaibshenjingvisitorpadfetchcodeverify 访客通过PAD提交访客码
// alibaba.ib.shenjing.visitor.pad.fetchcodeverify
//
// 访客通过PAD提交访客码，录脸进入园区。
func Alibabaibshenjingvisitorpadfetchcodeverify(clt *core.SDKClient, req *shenjing.AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest, session string) (*shenjing.AlibabaibshenjingvisitorpadfetchcodeverifyAPIResponse, error) {
	var resp shenjing.AlibabaibshenjingvisitorpadfetchcodeverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
