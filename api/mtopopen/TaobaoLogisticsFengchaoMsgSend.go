package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaologisticsfengchaomsgsend 丰巢走淘宝发包裹状态通知接口
// taobao.logistics.fengchao.msg.send
//
// 丰巢走淘宝发包裹状态通知接口
func Taobaologisticsfengchaomsgsend(clt *core.SDKClient, req *mtopopen.TaobaologisticsfengchaomsgsendAPIRequest, session string) (*mtopopen.TaobaologisticsfengchaomsgsendAPIResponse, error) {
	var resp mtopopen.TaobaologisticsfengchaomsgsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
