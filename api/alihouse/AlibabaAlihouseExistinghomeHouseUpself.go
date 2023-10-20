package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseUpself 房源上架
// alibaba.alihouse.existinghome.house.upself
//
// 房源信息上架
func AlibabaAlihouseExistinghomeHouseUpself(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseUpselfAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseUpselfAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
