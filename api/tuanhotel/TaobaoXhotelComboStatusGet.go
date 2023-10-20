package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// Taobaoxhotelcombostatusget 酒店宝贝状态查询
// taobao.xhotel.combo.status.get
//
// 酒店宝贝状态查询
func Taobaoxhotelcombostatusget(clt *core.SDKClient, req *tuanhotel.TaobaoxhotelcombostatusgetAPIRequest, session string) (*tuanhotel.TaobaoxhotelcombostatusgetAPIResponse, error) {
	var resp tuanhotel.TaobaoxhotelcombostatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
