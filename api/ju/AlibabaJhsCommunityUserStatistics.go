package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// Alibabajhscommunityuserstatistics 聚划算社群用户行为上报
// alibaba.jhs.community.user.statistics
//
// 聚划算社群用户行为上报
func Alibabajhscommunityuserstatistics(clt *core.SDKClient, req *ju.AlibabajhscommunityuserstatisticsAPIRequest, session string) (*ju.AlibabajhscommunityuserstatisticsAPIResponse, error) {
	var resp ju.AlibabajhscommunityuserstatisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
