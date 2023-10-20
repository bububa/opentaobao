package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiItempropsGet 拍卖相关类目属性
// taobao.paimai.itemprops.get
//
// 读取拍卖相关类目属性
func TaobaoPaimaiItempropsGet(clt *core.SDKClient, req *paimai.TaobaoPaimaiItempropsGetAPIRequest, resp *paimai.TaobaoPaimaiItempropsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
