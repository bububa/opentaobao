package caipiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取可用的彩种列表 APIResponse
taobao.caipiao.lotterytypes.get

获取彩票系统支持的可用于赠送的彩种列表
*/
type TaobaoCaipiaoLotterytypesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCaipiaoLotterytypesGetResponse `json:"caipiao_lotterytypes_get_response,omitempty"` 
    TaobaoCaipiaoLotterytypesGetResponse
}

/* model for simplify = false
type TaobaoCaipiaoLotterytypesGetResponse struct {

    // 彩种个数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 彩种的数据结构
    
    Results  struct {
        LotteryType  []LotteryType `json:"lottery_type,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoCaipiaoLotterytypesGetResponse struct {

    // 彩种个数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 彩种的数据结构
    Results   []LotteryType `json:"results,omitempty"`

}
