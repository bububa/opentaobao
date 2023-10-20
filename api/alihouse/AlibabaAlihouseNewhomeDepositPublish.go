package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeDepositPublish 创建、修改、发布房产预存金商品
// alibaba.alihouse.newhome.deposit.publish
//
// 创建、修改、发布房产预存金商品
func AlibabaAlihouseNewhomeDepositPublish(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeDepositPublishAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeDepositPublishAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
