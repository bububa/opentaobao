package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappwidgettemplateinstanceupdate 小部件实例化版本更新
// taobao.miniapp.widget.template.instance.update
//
// 小部件版本更新
func Taobaominiappwidgettemplateinstanceupdate(clt *core.SDKClient, req *miniappopen.TaobaominiappwidgettemplateinstanceupdateAPIRequest, session string) (*miniappopen.TaobaominiappwidgettemplateinstanceupdateAPIResponse, error) {
	var resp miniappopen.TaobaominiappwidgettemplateinstanceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
