package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfciotqrypersoninfo 查询物联卡个人实人认证信息
// alibaba.aliqin.fc.iot.qry.personinfo
//
// 查询物联卡个人实人认证信息
func Alibabaaliqinfciotqrypersoninfo(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotqrypersoninfoAPIRequest, session string) (*aliqin.AlibabaaliqinfciotqrypersoninfoAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotqrypersoninfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
