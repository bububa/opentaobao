package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaIworkCoreHrsGetperson 获取神鲸用户基本信息
// alibaba.iwork.core.hrs.getperson
//
// 神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询
func AlibabaIworkCoreHrsGetperson(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaIworkCoreHrsGetpersonAPIRequest, resp *campus.AlibabaIworkCoreHrsGetpersonAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
