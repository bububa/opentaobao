package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseDownself 房源信息下架
// alibaba.alihouse.existinghome.house.downself
//
// 房源信息下架
func AlibabaAlihouseExistinghomeHouseDownself(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseDownselfAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseDownselfAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
