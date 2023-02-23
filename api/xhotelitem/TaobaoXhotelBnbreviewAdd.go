package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbreviewAdd 对外开放评论接口
// taobao.xhotel.bnbreview.add
//
// 对外开放评论接口
func TaobaoXhotelBnbreviewAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbreviewAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbreviewAddAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbreviewAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
