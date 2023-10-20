package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaopaimaiitempropsget 拍卖相关类目属性
// taobao.paimai.itemprops.get
//
// 读取拍卖相关类目属性
func Taobaopaimaiitempropsget(clt *core.SDKClient, req *paimai.TaobaopaimaiitempropsgetAPIRequest, session string) (*paimai.TaobaopaimaiitempropsgetAPIResponse, error) {
	var resp paimai.TaobaopaimaiitempropsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
