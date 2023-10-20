package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaologisticstaowaimsgsend 淘外包裹物流信息走淘宝发包裹状态通知接口
// taobao.logistics.taowai.msg.send
//
// 淘外包裹物流信息走淘宝发包裹状态通知接口
func Taobaologisticstaowaimsgsend(clt *core.SDKClient, req *mtopopen.TaobaologisticstaowaimsgsendAPIRequest, session string) (*mtopopen.TaobaologisticstaowaimsgsendAPIResponse, error) {
	var resp mtopopen.TaobaologisticstaowaimsgsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
