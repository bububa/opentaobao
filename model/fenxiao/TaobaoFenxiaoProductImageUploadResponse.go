package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品图片上传 APIResponse
taobao.fenxiao.product.image.upload

产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
*/
type TaobaoFenxiaoProductImageUploadAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductImageUploadResponse
}

type TaobaoFenxiaoProductImageUploadResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_image_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 操作时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`

    
}
