package caipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

/* TaobaoCaipiaoSignstatusCheck
检查用户是否签署支付宝代购协议
taobao.caipiao.signstatus.check

检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。 */
func TaobaoCaipiaoSignstatusCheck(clt *core.SDKClient, req *caipiao.TaobaoCaipiaoSignstatusCheckAPIRequest, session string) (*caipiao.TaobaoCaipiaoSignstatusCheckAPIResponse, error) {
	var resp caipiao.TaobaoCaipiaoSignstatusCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
