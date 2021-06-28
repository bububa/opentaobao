package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询ODPS服务 APIResponse
odps.aliyuncs.com.DescribeOdpsService.2014-12-15

查询ODPS服务
*/
type OdpsAliyuncsComDescribeOdpsService2014-12-15APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"odps_aliyuncs_com_DescribeOdpsService_2014-12-15_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求ID
    
    RequestId   string `json:"RequestId,omitempty" xml:"