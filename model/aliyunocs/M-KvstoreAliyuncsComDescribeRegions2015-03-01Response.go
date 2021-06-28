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
    M-KvstoreAliyuncsComDescribeRegions2015-03-01Response
}

type M-KvstoreAliyuncsComDescribeRegions2015-03-01Response struct {
    XMLName xml.Name `xml:"m-kvstore_aliyuncs_com_DescribeRegions_2015-03-01_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求的唯一ID
    
    RequestId   string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`

    
    // RegionIds 为一个 List,里面每个元素由三部分组成【RegionId,LocalName,ZoneIds】，其中ZoneIds是多个值组成的String，多个值间用英文半角“,”分隔。
    
    RegionIds   []AliyunOcsRegion `json:"RegionIds,omitempty" xml:"RegionIds>aliyun_ocs_region,omitempty"`
    
    
}
