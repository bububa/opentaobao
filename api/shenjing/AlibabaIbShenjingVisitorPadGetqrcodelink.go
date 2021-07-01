package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

/* AlibabaIbShenjingVisitorPadGetqrcodelink
pad获取二维码
alibaba.ib.shenjing.visitor.pad.getqrcodelink

pad获取二维码链接。扫码录入人脸。 */
func AlibabaIbShenjingVisitorPadGetqrcodelink(clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest, session string) (*shenjing.AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse, error) {
	var resp shenjing.AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
