package aliyunocs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查看Region列表 APIResponse
m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01

查看Region列表
*/
type M-KvstoreAliyuncsComDescribeRegions2015-03-01APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"m-kvstore_aliyuncs_com_DescribeRegions_2015-03-01_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求的唯一ID
    
    RequestId   string `json:"RequestId,omitempty" xml:"