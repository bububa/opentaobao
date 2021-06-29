package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询用户全量的门店域商品接口(每页最多20条) API返回值 
alibaba.lst.pos.open.goods.getgoodsbypaging

分页查询用户全量的门店域商品接口(每页最多20条)
*/
type AlibabaLstPosOpenGoodsGetgoodsbypagingAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenGoodsGetgoodsbypagingResponse
}

// 分页查询用户全量的门店域商品接口(每页最多20条) 成功返回结果
type AlibabaLstPosOpenGoodsGetgoodsbypagingResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_goods_getgoodsbypaging_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *AlibabaLstPosOpenGoodsGetgoodsbypagingResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
