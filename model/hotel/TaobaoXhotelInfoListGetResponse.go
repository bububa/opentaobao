package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
酒店详细信息查询 APIResponse
taobao.xhotel.info.list.get

获取酒店详情信息
*/
type TaobaoXhotelInfoListGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoXhotelInfoListGetResponse `json:"taobao_xhotel_info_list_get_response,omitempty"`
}

type TaobaoXhotelInfoListGetResponse struct {

    // 酒店总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 标准酒店信息
    Hotels   []SHotelInfoObject `json:"hotels,omitempty"`

}
