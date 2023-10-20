package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizbrandinsert 新增电子价签商家
// taobao.uscesl.biz.brand.insert
//
// 一个电子价签业务身份下新增商家接口
func Taobaousceslbizbrandinsert(clt *core.SDKClient, req *uscesl.TaobaousceslbizbrandinsertAPIRequest, session string) (*uscesl.TaobaousceslbizbrandinsertAPIResponse, error) {
	var resp uscesl.TaobaousceslbizbrandinsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
