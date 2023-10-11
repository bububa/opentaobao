package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityActivityList 聚划算用增淘外社群服务活动列表
// alibaba.jhs.community.activity.list
//
// 聚划算用增淘外社群服务活动列表
func AlibabaJhsCommunityActivityList(clt *core.SDKClient, req *ju.AlibabaJhsCommunityActivityListAPIRequest, session string) (*ju.AlibabaJhsCommunityActivityListAPIResponse, error) {
	var resp ju.AlibabaJhsCommunityActivityListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
