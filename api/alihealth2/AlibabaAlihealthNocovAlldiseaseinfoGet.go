package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthNocovAlldiseaseinfoGet 获取全国疫情统计数据
// alibaba.alihealth.nocov.alldiseaseinfo.get
//
// 获取全国疫情统计数据
func AlibabaAlihealthNocovAlldiseaseinfoGet(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest, resp *alihealth2.AlibabaAlihealthNocovAlldiseaseinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
