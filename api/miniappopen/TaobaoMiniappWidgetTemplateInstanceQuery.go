package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappwidgettemplateinstancequery 小部件实例化版本查询
// taobao.miniapp.widget.template.instance.query
//
// 小部件实例化版本查询
func Taobaominiappwidgettemplateinstancequery(clt *core.SDKClient, req *miniappopen.TaobaominiappwidgettemplateinstancequeryAPIRequest, session string) (*miniappopen.TaobaominiappwidgettemplateinstancequeryAPIResponse, error) {
	var resp miniappopen.TaobaominiappwidgettemplateinstancequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
