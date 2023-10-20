package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoKaolaScitemAdd 考拉货品新增接口
// taobao.kaola.scitem.add
//
// 考拉货品新增接口
func TaobaoKaolaScitemAdd(clt *core.SDKClient, req *fenxiao.TaobaoKaolaScitemAddAPIRequest, resp *fenxiao.TaobaoKaolaScitemAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
