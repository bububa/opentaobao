package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoVasServiceGetServTimes 查询某个用户图片空间的使用情况
// taobao.vas.service.getServTimes
//
// 查询某个用户图片空间的使用情况
func TaobaoVasServiceGetServTimes(ctx context.Context, clt *core.SDKClient, req *media.TaobaoVasServiceGetServTimesAPIRequest, resp *media.TaobaoVasServiceGetServTimesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
