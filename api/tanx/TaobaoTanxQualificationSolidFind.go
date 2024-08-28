package tanx

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxQualificationSolidFind 客户固态共享资质
// taobao.tanx.qualification.solid.find
//
// 接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。
func TaobaoTanxQualificationSolidFind(ctx context.Context, clt *core.SDKClient, req *tanx.TaobaoTanxQualificationSolidFindAPIRequest, resp *tanx.TaobaoTanxQualificationSolidFindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
