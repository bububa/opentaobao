package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Taobaoalitriptravelnormalvisagetcompany 获取物流公司信息
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
func Taobaoalitriptravelnormalvisagetcompany(clt *core.SDKClient, req *normalvisa.TaobaoalitriptravelnormalvisagetcompanyAPIRequest, session string) (*normalvisa.TaobaoalitriptravelnormalvisagetcompanyAPIResponse, error) {
	var resp normalvisa.TaobaoalitriptravelnormalvisagetcompanyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
