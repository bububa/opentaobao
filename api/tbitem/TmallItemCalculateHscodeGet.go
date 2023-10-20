package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemcalculatehscodeget 算法获取hscode
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
func Tmallitemcalculatehscodeget(clt *core.SDKClient, req *tbitem.TmallitemcalculatehscodegetAPIRequest, session string) (*tbitem.TmallitemcalculatehscodegetAPIResponse, error) {
	var resp tbitem.TmallitemcalculatehscodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
