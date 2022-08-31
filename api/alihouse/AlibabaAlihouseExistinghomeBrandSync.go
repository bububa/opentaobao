package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrandSync 二手房公司品牌数据同步
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
func AlibabaAlihouseExistinghomeBrandSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrandSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeBrandSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeBrandSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
