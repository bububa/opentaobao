package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusCoreCompanycampusGetcombycamid 根据园区ID获取运营公司信息
// alibaba.campus.core.companycampus.getcombycamid
//
// 根据园区ID获取运营公司信息
func AlibabaCampusCoreCompanycampusGetcombycamid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest, resp *campus.AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
