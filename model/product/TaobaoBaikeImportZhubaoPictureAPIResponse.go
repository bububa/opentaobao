package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百科图片数据导入 API返回值 
taobao.baike.import.zhubao.picture

用于接入外部--图片--录入到商品百科中
*/
type TaobaoBaikeImportZhubaoPictureAPIResponse struct {
    model.CommonResponse
    TaobaoBaikeImportZhubaoPictureAPIResponseModel
}

// 百科图片数据导入 成功返回结果
type TaobaoBaikeImportZhubaoPictureAPIResponseModel struct {
    XMLName xml.Name `xml:"baike_import_zhubao_picture_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoBaikeImportZhubaoPictureResult `json:"result,omitempty" xml:"result,omitempty"`
}
