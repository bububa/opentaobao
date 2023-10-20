package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopconnectoreventpublish 连接器事件发布
// taobao.top.connector.event.publish
//
// 连接器事件发布
func Taobaotopconnectoreventpublish(clt *core.SDKClient, req *util.TaobaotopconnectoreventpublishAPIRequest, session string) (*util.TaobaotopconnectoreventpublishAPIResponse, error) {
	var resp util.TaobaotopconnectoreventpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
