package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Taobaoalitriptravelnormalvisaget 获取签证记录
// taobao.alitrip.travel.normalvisa.get
//
// 用于获取普通签证的记录信息
func Taobaoalitriptravelnormalvisaget(clt *core.SDKClient, req *normalvisa.TaobaoalitriptravelnormalvisagetAPIRequest, session string) (*normalvisa.TaobaoalitriptravelnormalvisagetAPIResponse, error) {
	var resp normalvisa.TaobaoalitriptravelnormalvisagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
