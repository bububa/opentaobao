package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Taobaoalitriptravelnormalvisastoreuser 代填办理人信息
// taobao.alitrip.travel.normalvisa.storeuser
//
// 卖家代填买家填写办理人信息
func Taobaoalitriptravelnormalvisastoreuser(clt *core.SDKClient, req *normalvisa.TaobaoalitriptravelnormalvisastoreuserAPIRequest, session string) (*normalvisa.TaobaoalitriptravelnormalvisastoreuserAPIResponse, error) {
	var resp normalvisa.TaobaoalitriptravelnormalvisastoreuserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
