package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugdownloadgetproductxml 获取product.xml的下载链接
// alibaba.alihealth.drug.download.getproductxml
//
// 阿里健康-追溯码-D2D获得药器信息下载地址，方便生产线操作
func Alibabaalihealthdrugdownloadgetproductxml(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugdownloadgetproductxmlAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugdownloadgetproductxmlAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugdownloadgetproductxmlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
