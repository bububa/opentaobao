package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

/* TaobaoQianniuCloudkefuOnlinestatuslogGet
查询客服在线状态
taobao.qianniu.cloudkefu.onlinestatuslog.get

按天查询客服账号的在线状态记录。如：登录，下线，挂起等
有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询 */
func TaobaoQianniuCloudkefuOnlinestatuslogGet(clt *core.SDKClient, req *qianniu.TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest, session string) (*qianniu.TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse, error) {
	var resp qianniu.TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
