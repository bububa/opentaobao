package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后信息同步 API返回值 
tmall.car.lease.postsynchronize

商家同步天猫开新车租后方案
*/
type TmallCarLeasePostsynchronizeAPIResponse struct {
    model.CommonResponse
    TmallCarLeasePostsynchronizeAPIResponseModel
}

// 天猫开新车租后信息同步 成功返回结果
type TmallCarLeasePostsynchronizeAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_car_lease_postsynchronize_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
