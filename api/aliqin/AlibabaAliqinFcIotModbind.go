package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfciotmodbind 物联网绑定/换绑API
// alibaba.aliqin.fc.iot.modbind
//
// 支持用户的设备的换绑和解绑操作
func Alibabaaliqinfciotmodbind(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotmodbindAPIRequest, session string) (*aliqin.AlibabaaliqinfciotmodbindAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotmodbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
