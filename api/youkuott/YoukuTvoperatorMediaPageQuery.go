package youkuott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// YoukuTvoperatorMediaPageQuery 运营商全量媒资分页查询
// youku.tvoperator.media.page.query
//
// 分页获取渠道全量媒资
func YoukuTvoperatorMediaPageQuery(clt *core.SDKClient, req *youkuott.YoukuTvoperatorMediaPageQueryAPIRequest, session string) (*youkuott.YoukuTvoperatorMediaPageQueryAPIResponse, error) {
	var resp youkuott.YoukuTvoperatorMediaPageQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
