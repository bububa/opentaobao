package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaouscesliteminfobatchput 批量写入商品信息接口
// taobao.uscesl.iteminfo.batch.put
//
// 电子架签批量写入商品数据，用于电子价签展示
func Taobaouscesliteminfobatchput(clt *core.SDKClient, req *uscesl.TaobaouscesliteminfobatchputAPIRequest, session string) (*uscesl.TaobaouscesliteminfobatchputAPIResponse, error) {
	var resp uscesl.TaobaouscesliteminfobatchputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
