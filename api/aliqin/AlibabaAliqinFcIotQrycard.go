package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfciotqrycard 查询终端信息
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
func Alibabaaliqinfciotqrycard(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotqrycardAPIRequest, session string) (*aliqin.AlibabaaliqinfciotqrycardAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotqrycardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
