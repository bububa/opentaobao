package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse 获取product.xml的下载链接 API返回值
// alibaba.alihealth.drug.download.getproductxml
//
// 阿里健康-追溯码-D2D获得药器信息下载地址，方便生产线操作
type AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugDownloadGetproductxmlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugDownloadGetproductxmlAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugDownloadGetproductxmlAPIResponseModel is 获取product.xml的下载链接 成功返回结果
type AlibabaAlihealthDrugDownloadGetproductxmlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_getproductxml_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataEntTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadGetproductxmlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugDownloadGetproductxmlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse)
	},
}

// GetAlibabaAlihealthDrugDownloadGetproductxmlAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse
func GetAlibabaAlihealthDrugDownloadGetproductxmlAPIResponse() *AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse {
	return poolAlibabaAlihealthDrugDownloadGetproductxmlAPIResponse.Get().(*AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse)
}

// ReleaseAlibabaAlihealthDrugDownloadGetproductxmlAPIResponse 将 AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadGetproductxmlAPIResponse(v *AlibabaAlihealthDrugDownloadGetproductxmlAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadGetproductxmlAPIResponse.Put(v)
}
