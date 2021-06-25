package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询ODPS服务 APIResponse
odps.aliyuncs.com.DescribeOdpsService.2014-12-15

查询ODPS服务
*/
type OdpsAliyuncsComDescribeOdpsService2014-12-15APIResponse struct {
    model.CommonResponse
    Response *OdpsAliyuncsComDescribeOdpsService2014-12-15Response `json:"odps_aliyuncs_com_DescribeOdpsService_2014-12-15_response,omitempty"`
}

type OdpsAliyuncsComDescribeOdpsService2014-12-15Response struct {

    // 请求ID
    RequestId   string `json:"RequestId,omitempty"`

    // 服务Id
    ServiceId   string `json:"ServiceId,omitempty"`

    // 开通服务时间，ISO 8601时间格式
    OpeningTime   string `json:"OpeningTime,omitempty"`

    // 业务锁定状态，例如：欠费，安全等。
    OperationLocks   string `json:"OperationLocks,omitempty"`

}
