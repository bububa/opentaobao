package lstlogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstLogisticsThirdpartCompanyList 供应商-异云-第三方物流公司列表
// alibaba.lst.logistics.thirdpart.company.list
//
// 异地云仓发货时，需填写的第三方物流公司列表
func AlibabaLstLogisticsThirdpartCompanyList(ctx context.Context, clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsThirdpartCompanyListAPIRequest, resp *lstlogistics.AlibabaLstLogisticsThirdpartCompanyListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
