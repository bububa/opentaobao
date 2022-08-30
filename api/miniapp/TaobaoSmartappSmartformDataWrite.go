package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappSmartformDataWrite 智能表单外部更新数据
// taobao.smartapp.smartform.data.write
//
// 智能表单外部更新数据
func TaobaoSmartappSmartformDataWrite(clt *core.SDKClient, req *miniapp.TaobaoSmartappSmartformDataWriteAPIRequest, session string) (*miniapp.TaobaoSmartappSmartformDataWriteAPIResponse, error) {
	var resp miniapp.TaobaoSmartappSmartformDataWriteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
