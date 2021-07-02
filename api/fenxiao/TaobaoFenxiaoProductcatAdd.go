package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatAdd 新增产品线
// taobao.fenxiao.productcat.add
//
// 新增产品线
func TaobaoFenxiaoProductcatAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatAddAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductcatAddAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductcatAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
