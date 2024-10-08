package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoMetaReceive 汽车说明书元数据上传
// tmall.aliauto.meta.receive
//
// 天猫汽车对外提供的汽车资源元数据上传接口
func TmallAliautoMetaReceive(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoMetaReceiveAPIRequest, resp *tmallcar.TmallAliautoMetaReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
