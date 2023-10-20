package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscadzonecreate 淘宝客-服务商-创建推广者位
// taobao.tbk.sc.adzone.create
//
// 提供淘宝客创建广告位
func Taobaotbkscadzonecreate(clt *core.SDKClient, req *tbk.TaobaotbkscadzonecreateAPIRequest, session string) (*tbk.TaobaotbkscadzonecreateAPIResponse, error) {
	var resp tbk.TaobaotbkscadzonecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
