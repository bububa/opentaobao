package lsticitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsticitem"
)

// Alibabalsticiteminfoquery 商品信息查询
// alibaba.lst.ic.item.info.query
//
// 查询商品信息
func Alibabalsticiteminfoquery(clt *core.SDKClient, req *lsticitem.AlibabalsticiteminfoqueryAPIRequest, session string) (*lsticitem.AlibabalsticiteminfoqueryAPIResponse, error) {
	var resp lsticitem.AlibabalsticiteminfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
