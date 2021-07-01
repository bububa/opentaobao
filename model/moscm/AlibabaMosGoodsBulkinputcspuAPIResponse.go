package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量录入商品信息 API返回值 
alibaba.mos.goods.bulkinputcspu

用于商品信息的批量导入到银泰商品中台
*/
type AlibabaMosGoodsBulkinputcspuAPIResponse struct {
    model.CommonResponse
    AlibabaMosGoodsBulkinputcspuAPIResponseModel
}

// 批量录入商品信息 成功返回结果
type AlibabaMosGoodsBulkinputcspuAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mos_goods_bulkinputcspu_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Data   *BulkInputCspuResult `json:"data,omitempty" xml:"data,omitempty"`
}
