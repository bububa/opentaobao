package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopeventpublish 同步事件发布
// taobao.top.event.publish
//
// 同步事件发布
func Taobaotopeventpublish(clt *core.SDKClient, req *util.TaobaotopeventpublishAPIRequest, session string) (*util.TaobaotopeventpublishAPIResponse, error) {
	var resp util.TaobaotopeventpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
