package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkspreadget 淘宝客-公用-长链转短链
// taobao.tbk.spread.get
//
// 输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
// 现阶段只支持短连接。
func Taobaotbkspreadget(clt *core.SDKClient, req *tbk.TaobaotbkspreadgetAPIRequest, session string) (*tbk.TaobaotbkspreadgetAPIResponse, error) {
	var resp tbk.TaobaotbkspreadgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
