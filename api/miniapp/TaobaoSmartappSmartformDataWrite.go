package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappSmartformDataWrite 智能表单外部更新数据
// taobao.smartapp.smartform.data.write
//
// 智能表单外部更新数据
func TaobaoSmartappSmartformDataWrite(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoSmartappSmartformDataWriteAPIRequest, resp *miniapp.TaobaoSmartappSmartformDataWriteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
