package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScUcrowdGet 淘宝客-服务商-获取人群标签
// taobao.tbk.sc.ucrowd.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群信息，包括人群名称描述及覆盖会员数。
func TaobaoTbkScUcrowdGet(clt *core.SDKClient, req *tbk.TaobaoTbkScUcrowdGetAPIRequest, session string) (*tbk.TaobaoTbkScUcrowdGetAPIResponse, error) {
	var resp tbk.TaobaoTbkScUcrowdGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
