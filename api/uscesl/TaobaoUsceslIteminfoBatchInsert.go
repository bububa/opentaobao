package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaouscesliteminfobatchinsert 按商家批量写入商品接口
// taobao.uscesl.iteminfo.batch.insert
//
// 【电子价签】支持按照商家-门店维度批量写入商品数据
func Taobaouscesliteminfobatchinsert(clt *core.SDKClient, req *uscesl.TaobaouscesliteminfobatchinsertAPIRequest, session string) (*uscesl.TaobaouscesliteminfobatchinsertAPIResponse, error) {
	var resp uscesl.TaobaouscesliteminfobatchinsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
