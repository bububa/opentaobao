package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscucrowdget 淘宝客-服务商-获取人群标签
// taobao.tbk.sc.ucrowd.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群信息，包括人群名称描述及覆盖会员数。
func Taobaotbkscucrowdget(clt *core.SDKClient, req *tbk.TaobaotbkscucrowdgetAPIRequest, session string) (*tbk.TaobaotbkscucrowdgetAPIResponse, error) {
	var resp tbk.TaobaotbkscucrowdgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
