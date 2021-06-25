package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户图片空间的使用情况 APIResponse
taobao.vas.service.getServTimes

查询某个用户图片空间的使用情况
*/
type TaobaoVasServiceGetServTimesAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVasServiceGetServTimesResponse `json:"taobao_vas_service_getServTimes_response,omitempty"`
}

type TaobaoVasServiceGetServTimesResponse struct {

    // 总次数（容量）
    TotalNum   int64 `json:"total_num,omitempty"`

    // 剩余次数（容量）
    LeftNum   int64 `json:"left_num,omitempty"`

}
