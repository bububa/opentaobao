package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectPresalepermitSubmit 提交预售证
// alibaba.alihouse.newhome.project.presalepermit.submit
//
// 提交楼盘预售证信息
func AlibabaAlihouseNewhomeProjectPresalepermitSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
