package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateHotproductDownload 联盟营销爆品下载接口
// aliexpress.affiliate.hotproduct.download
//
// 查询联盟爆品API
func AliexpressAffiliateHotproductDownload(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateHotproductDownloadAPIRequest, resp *aecreatives.AliexpressAffiliateHotproductDownloadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
