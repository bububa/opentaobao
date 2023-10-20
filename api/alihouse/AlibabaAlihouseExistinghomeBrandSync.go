package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrandSync 二手房公司品牌数据同步
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
func AlibabaAlihouseExistinghomeBrandSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrandSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrandSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
