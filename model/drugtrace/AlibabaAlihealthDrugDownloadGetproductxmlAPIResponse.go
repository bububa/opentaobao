package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloadgetproductxmlAPIResponse 获取product.xml的下载链接 API返回值
// alibaba.alihealth.drug.download.getproductxml
//
// 阿里健康-追溯码-D2D获得药器信息下载地址，方便生产线操作
type AlibabaalihealthdrugdownloadgetproductxmlAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugdownloadgetproductxmlAPIResponseModel
}

// AlibabaalihealthdrugdownloadgetproductxmlAPIResponseModel is 获取product.xml的下载链接 成功返回结果
type AlibabaalihealthdrugdownloadgetproductxmlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_getproductxml_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataEntTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
