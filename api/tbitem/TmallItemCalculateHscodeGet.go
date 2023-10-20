package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemCalculateHscodeGet 算法获取hscode
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
func TmallItemCalculateHscodeGet(clt *core.SDKClient, req *tbitem.TmallItemCalculateHscodeGetAPIRequest, resp *tbitem.TmallItemCalculateHscodeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
