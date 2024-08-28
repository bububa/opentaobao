package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliyunindepDigitalsmsCreatetemplate 数字短信模板创建
// alibaba.aliyunindep.digitalsms.createtemplate
//
// 数字短信模板创建，给阿里云一方产品使用，类型：9
func AlibabaAliyunindepDigitalsmsCreatetemplate(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest, resp *aliqin.AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
