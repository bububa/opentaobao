package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStorecategoryGet 门店类目获取接口
// taobao.qimen.storecategory.get
//
// 商家在ERP中调用该接口，获取门店类目
func TaobaoQimenStorecategoryGet(clt *core.SDKClient, req *qimen.TaobaoQimenStorecategoryGetAPIRequest, resp *qimen.TaobaoQimenStorecategoryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
