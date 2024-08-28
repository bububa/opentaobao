package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeCompanySync 二手房公司同步接口
// alibaba.alihouse.existinghome.company.sync
//
// 二手房公司同步接口
func AlibabaAlihouseExistinghomeCompanySync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeCompanySyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeCompanySyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
