package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaowirelessbuntingshopshorturlcreate 通过店铺id取得短链
// taobao.wireless.bunting.shop.shorturl.create
//
// 通过店铺id取得短链
func Taobaowirelessbuntingshopshorturlcreate(clt *core.SDKClient, req *mtopopen.TaobaowirelessbuntingshopshorturlcreateAPIRequest, session string) (*mtopopen.TaobaowirelessbuntingshopshorturlcreateAPIResponse, error) {
	var resp mtopopen.TaobaowirelessbuntingshopshorturlcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
