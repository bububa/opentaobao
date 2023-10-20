package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaohttpdnsget TOPDNS配置
// taobao.httpdns.get
//
// 获取TOP DNS配置
func Taobaohttpdnsget(clt *core.SDKClient, req *util.TaobaohttpdnsgetAPIRequest, session string) (*util.TaobaohttpdnsgetAPIResponse, error) {
	var resp util.TaobaohttpdnsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
