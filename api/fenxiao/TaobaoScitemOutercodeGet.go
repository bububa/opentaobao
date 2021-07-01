package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoScitemOutercodeGet
根据outerCode查询商品
taobao.scitem.outercode.get

根据outerCode查询商品 */
func TaobaoScitemOutercodeGet(clt *core.SDKClient, req *fenxiao.TaobaoScitemOutercodeGetAPIRequest, session string) (*fenxiao.TaobaoScitemOutercodeGetAPIResponse, error) {
	var resp fenxiao.TaobaoScitemOutercodeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
