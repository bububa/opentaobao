package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoscitemquery 查询后端商品
// taobao.scitem.query
//
// 查询后端商品
func Taobaoscitemquery(clt *core.SDKClient, req *fenxiao.TaobaoscitemqueryAPIRequest, session string) (*fenxiao.TaobaoscitemqueryAPIResponse, error) {
	var resp fenxiao.TaobaoscitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
