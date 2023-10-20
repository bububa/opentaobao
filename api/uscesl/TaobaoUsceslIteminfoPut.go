package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaouscesliteminfoput 电子价签显示用商品信息写入
// taobao.uscesl.iteminfo.put
//
// 用于电子价签上显示的商品信息的写入，包含价格及促销信息
func Taobaouscesliteminfoput(clt *core.SDKClient, req *uscesl.TaobaouscesliteminfoputAPIRequest, session string) (*uscesl.TaobaouscesliteminfoputAPIResponse, error) {
	var resp uscesl.TaobaouscesliteminfoputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
