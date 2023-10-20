package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomedepositpublish 创建、修改、发布房产预存金商品
// alibaba.alihouse.newhome.deposit.publish
//
// 创建、修改、发布房产预存金商品
func Alibabaalihousenewhomedepositpublish(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomedepositpublishAPIRequest, session string) (*alihouse.AlibabaalihousenewhomedepositpublishAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomedepositpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
