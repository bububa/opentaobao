package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// Alibabajhscommunitysubmittingtext 聚划算社群动态文案下发接口
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
func Alibabajhscommunitysubmittingtext(clt *core.SDKClient, req *ju.AlibabajhscommunitysubmittingtextAPIRequest, session string) (*ju.AlibabajhscommunitysubmittingtextAPIResponse, error) {
	var resp ju.AlibabajhscommunitysubmittingtextAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
