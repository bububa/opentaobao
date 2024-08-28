package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectCooperationSubmit 提交KA合作楼盘
// alibaba.alihouse.newhome.project.cooperation.submit
//
// 提交KA合作楼盘
func AlibabaAlihouseNewhomeProjectCooperationSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
