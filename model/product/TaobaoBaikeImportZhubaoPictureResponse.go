package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百科图片数据导入 APIResponse
taobao.baike.import.zhubao.picture

用于接入外部--图片--录入到商品百科中
*/
type TaobaoBaikeImportZhubaoPictureAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaikeImportZhubaoPictureResponse `json:"taobao_baike_import_zhubao_picture_response,omitempty"`
}

type TaobaoBaikeImportZhubaoPictureResponse struct {

    // result
    Result   *TaobaoBaikeImportZhubaoPictureResult `json:"result,omitempty"`

}
