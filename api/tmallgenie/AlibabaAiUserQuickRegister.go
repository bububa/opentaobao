package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaiuserquickregister 精灵用户注册申请
// alibaba.ai.user.quick.register
//
// 人工智能实验室精灵用户注册申请接口，开放给Iot厂商做厂商会员数据上报
func Alibabaaiuserquickregister(clt *core.SDKClient, req *tmallgenie.AlibabaaiuserquickregisterAPIRequest, session string) (*tmallgenie.AlibabaaiuserquickregisterAPIResponse, error) {
	var resp tmallgenie.AlibabaaiuserquickregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
