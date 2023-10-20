package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizapdelete 删除价签AP设备
// taobao.uscesl.biz.ap.delete
//
// 删除价签AP设备
func Taobaousceslbizapdelete(clt *core.SDKClient, req *uscesl.TaobaousceslbizapdeleteAPIRequest, session string) (*uscesl.TaobaousceslbizapdeleteAPIResponse, error) {
	var resp uscesl.TaobaousceslbizapdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
