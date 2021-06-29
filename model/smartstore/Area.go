package smartstore

// Area 
type Area struct {
    // 区域类型.area区域 1:country/国家;2:province/省/自治区/直辖市;3:city/地区(省下面的地级市);4:district/县/市(县级市)/区;abroad:海外. 比如北京市的area_type = 2,朝阳区是北京市的一个区,所以朝阳区的area_type = 4.
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 地域名称.如北京市,杭州市,西湖区,每一个area_id 都代表了一个具体的地区.
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 父节点区域标识.如北京市的area_id是110100,朝阳区是北京市的一个区,所以朝阳区的parent_id就是北京市的area_id.
    ParentId   int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
    // 具体一个地区的邮编
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
    // 标准行政区域代码.参考:http://www.stats.gov.cn/tjbz/xzqhdm/t20120105_402777427.htm
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
}
