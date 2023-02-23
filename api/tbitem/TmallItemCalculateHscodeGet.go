package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemCalculateHscodeGet 算法获取hscode
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
func TmallItemCalculateHscodeGet(clt *core.SDKClient, req *tbitem.TmallItemCalculateHscodeGetAPIRequest, session string) (*tbitem.TmallItemCalculateHscodeGetAPIResponse, error) {
	var resp tbitem.TmallItemCalculateHscodeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
