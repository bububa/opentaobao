package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMsAreaProvinceList 疫苗预约省份列表查询
// alibaba.alihealth.ms.area.province.list
//
// 微信小程序疫苗预约省份列表查询
func AlibabaAlihealthMsAreaProvinceList(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMsAreaProvinceListAPIRequest, resp *alihealth2.AlibabaAlihealthMsAreaProvinceListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
