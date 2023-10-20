package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappSmartformDataWrite 智能表单外部更新数据
// taobao.smartapp.smartform.data.write
//
// 智能表单外部更新数据
func TaobaoSmartappSmartformDataWrite(clt *core.SDKClient, req *miniapp.TaobaoSmartappSmartformDataWriteAPIRequest, resp *miniapp.TaobaoSmartappSmartformDataWriteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
