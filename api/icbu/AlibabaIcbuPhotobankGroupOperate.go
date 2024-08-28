package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuPhotobankGroupOperate 图片银行分组操作接口
// alibaba.icbu.photobank.group.operate
//
// 修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
func AlibabaIcbuPhotobankGroupOperate(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankGroupOperateAPIRequest, resp *icbu.AlibabaIcbuPhotobankGroupOperateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
