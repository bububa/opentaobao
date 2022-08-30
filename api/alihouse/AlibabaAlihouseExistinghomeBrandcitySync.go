package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrandcitySync 二手房城市品牌店数据同步
// alibaba.alihouse.existinghome.brandcity.sync
//
// 二手房城市品牌店数据同步
func AlibabaAlihouseExistinghomeBrandcitySync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
