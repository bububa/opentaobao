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
type OdpsAliyuncsComDescribeOdpsService2014_12_15APIResponse struct {
    model.CommonResponse
    OdpsAliyuncsComDescribeOdpsService2014_12_15Response
}

type OdpsAliyuncsComDescribeOdpsService2014_12_15Response struct {
    XMLName xml.Name `xml:"odps_aliyuncs_com_DescribeOdpsService_2014-12-15_response"`
    

    // 请求ID
    
    RequestId   string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`

    
    // 服务Id
    
    ServiceId   string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`

    
    // 开通服务时间，ISO 8601时间格式
    
    OpeningTime   string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`

    
    // 业务锁定状态，例如：欠费，安全等。
    
    OperationLocks   string `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty"`

    
}
