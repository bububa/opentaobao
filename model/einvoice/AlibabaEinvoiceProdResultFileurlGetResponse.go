package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-发票文件下载地址查询 API返回值 
alibaba.einvoice.prod.result.fileurl.get

发票文件下载地址查询，外部ISV通过该接口可以查对应发票文件
*/
type AlibabaEinvoiceProdResultFileurlGetAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceProdResultFileurlGetResponse
}

// 发票中台-发票文件下载地址查询 成功返回结果
type AlibabaEinvoiceProdResultFileurlGetResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_prod_result_fileurl_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 发票版式文件下载地址。下载链接是一次性链接，有效期1天。请发票文件下载后本地保存， 若异常导致需要再次下载文件，请再次请求接口获取。
    FileDownloadUrl   string `json:"file_download_url,omitempty" xml:"file_download_url,omitempty"`
}
