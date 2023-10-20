package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliatehotproductdownload 联盟营销爆品下载接口
// aliexpress.affiliate.hotproduct.download
//
// 查询联盟爆品API
func Aliexpressaffiliatehotproductdownload(clt *core.SDKClient, req *aecreatives.AliexpressaffiliatehotproductdownloadAPIRequest, session string) (*aecreatives.AliexpressaffiliatehotproductdownloadAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliatehotproductdownloadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
