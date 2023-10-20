package yunos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunos"
)

// Yunoscosmodatapush COSMO-PUSH模式数据接入
// yunos.cosmo.data.push
//
// YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
func Yunoscosmodatapush(clt *core.SDKClient, req *yunos.YunoscosmodatapushAPIRequest, session string) (*yunos.YunoscosmodatapushAPIResponse, error) {
	var resp yunos.YunoscosmodatapushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
