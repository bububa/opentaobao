package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeRegionInfoSubmit 商圈专家信息同步
// alibaba.alihouse.existinghome.region.info.submit
//
// 商圈专家信息同步
func AlibabaAlihouseExistinghomeRegionInfoSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
