package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeDepositPublish 创建、修改、发布房产预存金商品
// alibaba.alihouse.newhome.deposit.publish
//
// 创建、修改、发布房产预存金商品
func AlibabaAlihouseNewhomeDepositPublish(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeDepositPublishAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeDepositPublishAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeDepositPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
