package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizapadd 新增价签通讯AP设备
// taobao.uscesl.biz.ap.add
//
// 根据门店和ap的MAC地址新增
func Taobaousceslbizapadd(clt *core.SDKClient, req *uscesl.TaobaousceslbizapaddAPIRequest, session string) (*uscesl.TaobaousceslbizapaddAPIResponse, error) {
	var resp uscesl.TaobaousceslbizapaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
