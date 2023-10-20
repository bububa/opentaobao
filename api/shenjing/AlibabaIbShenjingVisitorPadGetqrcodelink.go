package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// Alibabaibshenjingvisitorpadgetqrcodelink pad获取二维码
// alibaba.ib.shenjing.visitor.pad.getqrcodelink
//
// pad获取二维码链接。扫码录入人脸。
func Alibabaibshenjingvisitorpadgetqrcodelink(clt *core.SDKClient, req *shenjing.AlibabaibshenjingvisitorpadgetqrcodelinkAPIRequest, session string) (*shenjing.AlibabaibshenjingvisitorpadgetqrcodelinkAPIResponse, error) {
	var resp shenjing.AlibabaibshenjingvisitorpadgetqrcodelinkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
