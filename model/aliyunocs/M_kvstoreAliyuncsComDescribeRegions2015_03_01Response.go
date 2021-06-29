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
type M_kvstoreAliyuncsComDescribeRegions2015_03_01APIResponse struct {
    model.CommonResponse
    M_kvstoreAliyuncsComDescribeRegions2015_03_01Response
}

type M_kvstoreAliyuncsComDescribeRegions2015_03_01Response struct {
    XMLName xml.Name `xml:"m-kvstore_aliyuncs_com_DescribeRegions_2015-03-01_response"`
    

    // 请求的唯一ID
    
    RequestId   string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`

    
    // RegionIds 为一个 List,里面每个元素由三部分组成【RegionId,LocalName,ZoneIds】，其中ZoneIds是多个值组成的String，多个值间用英文半角“,”分隔。
    
    RegionIds   []AliyunOcsRegion `json:"RegionIds,omitempty" xml:"RegionIds>aliyun_ocs_region,omitempty"`
    
    
}
