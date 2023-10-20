package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// Alibabajhscommunitywechatlogin 聚划算用增淘外社群登录
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
func Alibabajhscommunitywechatlogin(clt *core.SDKClient, req *ju.AlibabajhscommunitywechatloginAPIRequest, session string) (*ju.AlibabajhscommunitywechatloginAPIResponse, error) {
	var resp ju.AlibabajhscommunitywechatloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
