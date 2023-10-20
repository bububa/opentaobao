package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiapptemplaterollback 回滚实例化应用
// taobao.miniapp.template.rollback
//
// 将实例化小程序回滚到指定版本
func Taobaominiapptemplaterollback(clt *core.SDKClient, req *miniappopen.TaobaominiapptemplaterollbackAPIRequest, session string) (*miniappopen.TaobaominiapptemplaterollbackAPIResponse, error) {
	var resp miniappopen.TaobaominiapptemplaterollbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
