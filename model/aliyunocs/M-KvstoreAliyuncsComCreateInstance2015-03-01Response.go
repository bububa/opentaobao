package aliyunocs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OCS实例 APIResponse
m-kvstore.aliyuncs.com.CreateInstance.2015-03-01

创建OCS实例
*/
type M-KvstoreAliyuncsComCreateInstance2015-03-01APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"m-kvstore_aliyuncs_com_CreateInstance_2015-03-01_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // OCS例ID;<br/>16位字符串(全局唯一)
    
    InstanceId   string `json:"InstanceId,omitempty" xml:"