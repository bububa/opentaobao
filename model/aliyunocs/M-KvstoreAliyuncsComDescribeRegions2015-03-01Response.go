package aliyunocs

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查看Region列表 APIResponse
m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01

查看Region列表
*/
type M-KvstoreAliyuncsComDescribeRegions2015-03-01APIResponse struct {
    model.CommonResponse
    // Response *M-KvstoreAliyuncsComDescribeRegions2015-03-01Response `json:"m-kvstore_aliyuncs_com_DescribeRegions_2015-03-01_response,omitempty"` 
    M-KvstoreAliyuncsComDescribeRegions2015-03-01Response
}

/* model for simplify = false
type M-KvstoreAliyuncsComDescribeRegions2015-03-01Response struct {

    // 请求的唯一ID
    
    RequestId   string `json:"RequestId,omitempty"`
    

    // RegionIds 为一个 List,里面每个元素由三部分组成【RegionId,LocalName,ZoneIds】，其中ZoneIds是多个值组成的String，多个值间用英文半角“,”分隔。
    
    RegionIds  struct {
        AliyunOcsRegion  []AliyunOcsRegion `json:"aliyun_ocs_region,omitempty"`
    } `json:"RegionIds,omitempty"`
    

}
*/

type M-KvstoreAliyuncsComDescribeRegions2015-03-01Response struct {

    // 请求的唯一ID
    RequestId   string `json:"RequestId,omitempty"`

    // RegionIds 为一个 List,里面每个元素由三部分组成【RegionId,LocalName,ZoneIds】，其中ZoneIds是多个值组成的String，多个值间用英文半角“,”分隔。
    RegionIds   []AliyunOcsRegion `json:"RegionIds,omitempty"`

}
