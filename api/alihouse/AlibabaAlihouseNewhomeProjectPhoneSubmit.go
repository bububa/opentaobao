package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectPhoneSubmit 提交楼盘电话
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
func AlibabaAlihouseNewhomeProjectPhoneSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
