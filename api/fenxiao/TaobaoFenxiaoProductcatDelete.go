package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatDelete 删除产品线
// taobao.fenxiao.productcat.delete
//
// 删除产品线
func TaobaoFenxiaoProductcatDelete(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatDeleteAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductcatDeleteAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductcatDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
