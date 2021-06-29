package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量录入商品信息 APIResponse
alibaba.mos.goods.bulkinputcspu

用于商品信息的批量导入到银泰商品中台
*/
type AlibabaMosGoodsBulkinputcspuAPIResponse struct {
    model.CommonResponse
    AlibabaMosGoodsBulkinputcspuResponse
}

type AlibabaMosGoodsBulkinputcspuResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_goods_bulkinputcspu_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回数据
    
    Data   *BulkInputCspuResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
