package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百科图片数据导入 APIResponse
taobao.baike.import.zhubao.picture

用于接入外部--图片--录入到商品百科中
*/
type TaobaoBaikeImportZhubaoPictureAPIResponse struct {
    model.CommonResponse
    TaobaoBaikeImportZhubaoPictureResponse
}

type TaobaoBaikeImportZhubaoPictureResponse struct {
    XMLName xml.Name `xml:"baike_import_zhubao_picture_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoBaikeImportZhubaoPictureResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
