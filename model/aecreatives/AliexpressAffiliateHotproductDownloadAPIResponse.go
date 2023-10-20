package aecreatives

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliatehotproductdownloadAPIResponse 联盟营销爆品下载接口 API返回值
// aliexpress.affiliate.hotproduct.download
//
// 查询联盟爆品API
type AliexpressaffiliatehotproductdownloadAPIResponse struct {
	model.CommonResponse
	AliexpressaffiliatehotproductdownloadAPIResponseModel
}

// AliexpressaffiliatehotproductdownloadAPIResponseModel is 联盟营销爆品下载接口 成功返回结果
type AliexpressaffiliatehotproductdownloadAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_hotproduct_download_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
