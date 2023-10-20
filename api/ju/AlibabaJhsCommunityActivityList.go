package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// Alibabajhscommunityactivitylist 聚划算用增淘外社群服务活动列表
// alibaba.jhs.community.activity.list
//
// 聚划算用增淘外社群服务活动列表
func Alibabajhscommunityactivitylist(clt *core.SDKClient, req *ju.AlibabajhscommunityactivitylistAPIRequest, session string) (*ju.AlibabajhscommunityactivitylistAPIResponse, error) {
	var resp ju.AlibabajhscommunityactivitylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
