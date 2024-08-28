package youkuott

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// YoukuTvoperatorMediaPageQuery 运营商全量媒资分页查询
// youku.tvoperator.media.page.query
//
// 分页获取渠道全量媒资
func YoukuTvoperatorMediaPageQuery(ctx context.Context, clt *core.SDKClient, req *youkuott.YoukuTvoperatorMediaPageQueryAPIRequest, resp *youkuott.YoukuTvoperatorMediaPageQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
