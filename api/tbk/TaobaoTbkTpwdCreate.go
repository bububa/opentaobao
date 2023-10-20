package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbktpwdcreate 淘宝客-公用-淘口令生成
// taobao.tbk.tpwd.create
//
// 提供淘口令生成接口。提交需要生成淘口令的内容、logo、url等参数，生成淘口令，其中关键信息为￥SADadW￥，后续可基于淘口令进行文案包装组装用于传播。
func Taobaotbktpwdcreate(clt *core.SDKClient, req *tbk.TaobaotbktpwdcreateAPIRequest, session string) (*tbk.TaobaotbktpwdcreateAPIResponse, error) {
	var resp tbk.TaobaotbktpwdcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
