package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoregionpricecancle 取消区域价格
// taobao.region.price.cancle
//
// 取消区域价格
func Taobaoregionpricecancle(clt *core.SDKClient, req *fenxiao.TaobaoregionpricecancleAPIRequest, session string) (*fenxiao.TaobaoregionpricecancleAPIResponse, error) {
	var resp fenxiao.TaobaoregionpricecancleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
