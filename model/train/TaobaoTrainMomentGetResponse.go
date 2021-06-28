package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
火车票时刻表 APIResponse
taobao.train.moment.get

查询火车票车次时刻表
*/
type TaobaoTrainMomentGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainMomentGetResponse `json:"train_moment_get_response,omitempty"` 
    TaobaoTrainMomentGetResponse
}

/* model for simplify = false
type TaobaoTrainMomentGetResponse struct {

    // 出发到达对应时刻表索引
    
    Bindex   string `json:"bindex,omitempty"`
    

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 经停站
    
    Stations  struct {
        StationPassInfo  []StationPassInfo `json:"station_pass_info,omitempty"`
    } `json:"stations,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoTrainMomentGetResponse struct {

    // 出发到达对应时刻表索引
    Bindex   string `json:"bindex,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误信息
    ResultMsg   string `json:"result_msg,omitempty"`

    // 经停站
    Stations   []StationPassInfo `json:"stations,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
