package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaaliqinfciotcardStatus 物联卡状态查询
// alibaba.aliqin.fc.iot.cardStatus
//
// 物联卡状态查询
func AlibabaaliqinfciotcardStatus(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotcardStatusAPIRequest, session string) (*aliqin.AlibabaaliqinfciotcardStatusAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotcardStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
