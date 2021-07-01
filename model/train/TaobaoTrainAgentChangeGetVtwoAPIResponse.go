package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取改签单详情v2--增加鉴权校验 API返回值 
taobao.train.agent.change.get.vtwo

卖家获取待处理的改签单详情
*/
type TaobaoTrainAgentChangeGetVtwoAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentChangeGetVtwoAPIResponseModel
}

// 获取改签单详情v2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentChangeGetVtwoAPIResponseModel struct {
    XMLName xml.Name `xml:"train_agent_change_get_vtwo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 扩展参数
    ExtendParam   string `json:"extend_param,omitempty" xml:"extend_param,omitempty"`
    // 1
    Tickets   []ChangeTicketInfo `json:"tickets,omitempty" xml:"tickets>change_ticket_info,omitempty"`
    // 坐席
    SeatName   int64 `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
    // 申请单
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 状态，1 待支付、2 待改签、3 已改签、4 改签失败
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 总改签费
    TotalChangeFee   int64 `json:"total_change_fee,omitempty" xml:"total_change_fee,omitempty"`
    // 车次
    TrainNum   string `json:"train_num,omitempty" xml:"train_num,omitempty"`
    // 出发站
    FromStationName   string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
    // 最晚改签时间
    LatestChangeTime   string `json:"latest_change_time,omitempty" xml:"latest_change_time,omitempty"`
    // 发车时间
    FromTime   string `json:"from_time,omitempty" xml:"from_time,omitempty"`
    // 主订单
    MainBizOrderId   int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
    // 到达时间
    ToTime   string `json:"to_time,omitempty" xml:"to_time,omitempty"`
    // 到达站
    ToStationName   string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
}
