package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// TaobaoXhotelComboStatusGet 酒店宝贝状态查询
// taobao.xhotel.combo.status.get
//
// 酒店宝贝状态查询
func TaobaoXhotelComboStatusGet(clt *core.SDKClient, req *tuanhotel.TaobaoXhotelComboStatusGetAPIRequest, session string) (*tuanhotel.TaobaoXhotelComboStatusGetAPIResponse, error) {
	var resp tuanhotel.TaobaoXhotelComboStatusGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
