package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscucrowdcreate 淘宝客-服务商-创建人群标签
// taobao.tbk.sc.ucrowd.create
//
// 服务商使用。可为淘宝客创建会员人群标签，获得人群id，人群id可用于物料评估推荐等场景。
func Taobaotbkscucrowdcreate(clt *core.SDKClient, req *tbk.TaobaotbkscucrowdcreateAPIRequest, session string) (*tbk.TaobaotbkscucrowdcreateAPIResponse, error) {
	var resp tbk.TaobaotbkscucrowdcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
