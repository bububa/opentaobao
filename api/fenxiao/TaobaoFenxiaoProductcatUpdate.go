package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductcatupdate 修改产品线
// taobao.fenxiao.productcat.update
//
// 修改产品线
func Taobaofenxiaoproductcatupdate(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductcatupdateAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductcatupdateAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductcatupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
