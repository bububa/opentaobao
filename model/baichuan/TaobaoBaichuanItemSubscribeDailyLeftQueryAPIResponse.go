package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询当天可添加的余量 API返回值 
taobao.baichuan.item.subscribe.daily.left.query

查询当天可添加的余量
*/
type TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponseModel
}

// 查询当天可添加的余量 成功返回结果
type TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"baichuan_item_subscribe_daily_left_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeDailyLeftQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
