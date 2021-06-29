package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品图片删除 APIResponse
taobao.fenxiao.product.image.delete

产品图片删除，只删除图片信息，不真正删除图片
*/
type TaobaoFenxiaoProductImageDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductImageDeleteResponse
}

type TaobaoFenxiaoProductImageDeleteResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_image_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 操作时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`

    
}
