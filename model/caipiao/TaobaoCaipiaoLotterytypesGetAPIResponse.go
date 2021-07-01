package caipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取可用的彩种列表 API返回值 
taobao.caipiao.lotterytypes.get

获取彩票系统支持的可用于赠送的彩种列表
*/
type TaobaoCaipiaoLotterytypesGetAPIResponse struct {
    model.CommonResponse
    TaobaoCaipiaoLotterytypesGetAPIResponseModel
}

// 获取可用的彩种列表 成功返回结果
type TaobaoCaipiaoLotterytypesGetAPIResponseModel struct {
    XMLName xml.Name `xml:"caipiao_lotterytypes_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 彩种个数
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
    // 彩种的数据结构
    Results   []LotteryType `json:"results,omitempty" xml:"results>lottery_type,omitempty"`
}
