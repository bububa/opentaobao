package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaWtCifCoopOsstokenGet 获取oss签名接口
// alibaba.wt.cif.coop.osstoken.get
//
// 商家合作上传oss图片获取token接口
func AlibabaWtCifCoopOsstokenGet(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaWtCifCoopOsstokenGetAPIRequest, resp *alicom.AlibabaWtCifCoopOsstokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
