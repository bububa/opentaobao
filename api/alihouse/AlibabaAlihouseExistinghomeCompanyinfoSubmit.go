package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeCompanyinfoSubmit 进件接口
// alibaba.alihouse.existinghome.companyinfo.submit
//
// 进件接口
func AlibabaAlihouseExistinghomeCompanyinfoSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
