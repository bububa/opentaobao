package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoKfcKeywordSearch 关键词过滤匹配
// taobao.kfc.keyword.search
//
// 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
func TaobaoKfcKeywordSearch(ctx context.Context, clt *core.SDKClient, req *util.TaobaoKfcKeywordSearchAPIRequest, resp *util.TaobaoKfcKeywordSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
