package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribelogsImport openim群聊天记录导入
// taobao.openim.tribelogs.import
//
// openim群聊天记录导入
func TaobaoOpenimTribelogsImport(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribelogsImportAPIRequest, resp *openim.TaobaoOpenimTribelogsImportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
