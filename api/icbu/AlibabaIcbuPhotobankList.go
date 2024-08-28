package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuPhotobankList 国际站图片银行查询接口
// alibaba.icbu.photobank.list
//
// 国际站图片银行查询接口
func AlibabaIcbuPhotobankList(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankListAPIRequest, resp *icbu.AlibabaIcbuPhotobankListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
