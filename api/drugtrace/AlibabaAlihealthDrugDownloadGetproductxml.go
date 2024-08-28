package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadGetproductxml 获取product.xml的下载链接
// alibaba.alihealth.drug.download.getproductxml
//
// 阿里健康-追溯码-D2D获得药器信息下载地址，方便生产线操作
func AlibabaAlihealthDrugDownloadGetproductxml(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest, resp *drugtrace.AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
