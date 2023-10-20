package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaosmartappsmartformdatawrite 智能表单外部更新数据
// taobao.smartapp.smartform.data.write
//
// 智能表单外部更新数据
func Taobaosmartappsmartformdatawrite(clt *core.SDKClient, req *miniapp.TaobaosmartappsmartformdatawriteAPIRequest, session string) (*miniapp.TaobaosmartappsmartformdatawriteAPIResponse, error) {
	var resp miniapp.TaobaosmartappsmartformdatawriteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
