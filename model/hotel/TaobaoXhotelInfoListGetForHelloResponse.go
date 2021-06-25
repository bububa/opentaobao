package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
哈罗获取酒店详情 APIResponse
taobao.xhotel.info.list.get.for.hello

哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
*/
type TaobaoXhotelInfoListGetForHelloAPIResponse struct {
    model.CommonResponse
    Response *TaobaoXhotelInfoListGetForHelloResponse `json:"taobao_xhotel_info_list_get_for_hello_response,omitempty"`
}

type TaobaoXhotelInfoListGetForHelloResponse struct {

    // 查询结果集
    Result   *TaobaoXhotelInfoListGetForHelloResultSet `json:"result,omitempty"`

}
