package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallMsfVerify 喵师傅核销接口
// tmall.msf.verify
//
// msf服务核销的top接口
func TmallMsfVerify(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallMsfVerifyAPIRequest, resp *servicecenter.TmallMsfVerifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
