package ioti

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// AlibabaItEslEslimageShowimagecommon 对混合云提供的刷图接口
// alibaba.it.esl.eslimage.showimagecommon
//
// 混合云使用，提供给isv和我们混合云环境部署的应用刷图
func AlibabaItEslEslimageShowimagecommon(ctx context.Context, clt *core.SDKClient, req *ioti.AlibabaItEslEslimageShowimagecommonAPIRequest, resp *ioti.AlibabaItEslEslimageShowimagecommonAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
