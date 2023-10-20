package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Taobaoalitriptravelnormalvisagetdetail 获取单笔订单的详情
// taobao.alitrip.travel.normalvisa.getdetail
//
// 获取单笔签证的详细记录
func Taobaoalitriptravelnormalvisagetdetail(clt *core.SDKClient, req *normalvisa.TaobaoalitriptravelnormalvisagetdetailAPIRequest, session string) (*normalvisa.TaobaoalitriptravelnormalvisagetdetailAPIResponse, error) {
	var resp normalvisa.TaobaoalitriptravelnormalvisagetdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
