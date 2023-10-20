package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfciotcardoffer 查询物联网卡上订购的offer
// alibaba.aliqin.fc.iot.cardoffer
//
// 查询物联网卡上订购的offer
func Alibabaaliqinfciotcardoffer(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotcardofferAPIRequest, session string) (*aliqin.AlibabaaliqinfciotcardofferAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotcardofferAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
