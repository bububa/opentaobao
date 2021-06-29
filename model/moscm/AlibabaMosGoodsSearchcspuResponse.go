package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
cspu查询 API返回值 
alibaba.mos.goods.searchcspu

商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流）
*/
type AlibabaMosGoodsSearchcspuAPIResponse struct {
    model.CommonResponse
    AlibabaMosGoodsSearchcspuResponse
}

// cspu查询 成功返回结果
type AlibabaMosGoodsSearchcspuResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_goods_searchcspu_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *PagedList `json:"data,omitempty" xml:"data,omitempty"`
}
