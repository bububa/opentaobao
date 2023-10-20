package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaaliqinfciotcardInfo 物联卡信息查询
// alibaba.aliqin.fc.iot.cardInfo
//
// 物联卡信息查询
func AlibabaaliqinfciotcardInfo(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotcardInfoAPIRequest, session string) (*aliqin.AlibabaaliqinfciotcardInfoAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotcardInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
