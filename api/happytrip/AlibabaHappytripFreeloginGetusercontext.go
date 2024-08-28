package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripFreeloginGetusercontext 提供给外部系统的免登校验
// alibaba.happytrip.freelogin.getusercontext
//
// 免登融合，提供免登相关接口给外部供应商做登录验证
func AlibabaHappytripFreeloginGetusercontext(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripFreeloginGetusercontextAPIRequest, resp *happytrip.AlibabaHappytripFreeloginGetusercontextAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
