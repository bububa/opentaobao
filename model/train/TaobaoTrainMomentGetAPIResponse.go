package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainMomentGetAPIResponse 火车票时刻表 API返回值
// taobao.train.moment.get
//
// 查询火车票车次时刻表
type TaobaoTrainMomentGetAPIResponse struct {
	model.CommonResponse
	TaobaoTrainMomentGetAPIResponseModel
}

// TaobaoTrainMomentGetAPIResponseModel is 火车票时刻表 成功返回结果
type TaobaoTrainMomentGetAPIResponseModel struct {
	XMLName xml.Name `xml:"train_moment_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 经停站
	Stations []StationPassInfo `json:"stations,omitempty" xml:"stations>station_pass_info,omitempty"`
	// 出发到达对应时刻表索引
	Bindex string `json:"bindex,omitempty" xml:"bindex,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
