package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqintcctradeidentityget 运营商获得用户身份信息
// alibaba.aliqin.tcc.trade.identity.get
//
// 天猫网厅运营商官方旗舰店获取用户身份信息
func Alibabaaliqintcctradeidentityget(clt *core.SDKClient, req *alicom.AlibabaaliqintcctradeidentitygetAPIRequest, session string) (*alicom.AlibabaaliqintcctradeidentitygetAPIResponse, error) {
	var resp alicom.AlibabaaliqintcctradeidentitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
