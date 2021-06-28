package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票时刻表 APIResponse
taobao.train.moment.get

查询火车票车次时刻表
*/
type TaobaoTrainMomentGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_moment_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出发到达对应时刻表索引
    
    Bindex   string `json:"bindex,omitempty" xml:"