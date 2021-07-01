package yunos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunos"
)

/* YunosCosmoDataPush
COSMO-PUSH模式数据接入
yunos.cosmo.data.push

YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入 */
func YunosCosmoDataPush(clt *core.SDKClient, req *yunos.YunosCosmoDataPushAPIRequest, session string) (*yunos.YunosCosmoDataPushAPIResponse, error) {
	var resp yunos.YunosCosmoDataPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
