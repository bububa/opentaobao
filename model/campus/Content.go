package campus

// Content 
type Content struct {
    // 面积
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 创建人
    Creator   string `json:"creator,omitempty" xml:"creator,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 排序号
    OrderNo   int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 类型
    SpaceType   string `json:"space_type,omitempty" xml:"space_type,omitempty"`
    // 园区名称
    CampusName   string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
    // 是否删除
    IsDelete   bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
    // 修改人
    Modifier   string `json:"modifier,omitempty" xml:"modifier,omitempty"`
    // 园区id
    CampusId   int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
    // 地图楼层id
    GeoFloorId   int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
    // 描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // uuid
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    // 楼宇id
    BuildingId   int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
    // 楼层id
    FloorId   int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
    // 楼宇名称
    BuildingName   string `json:"building_name,omitempty" xml:"building_name,omitempty"`
    // 公司id
    CompanyId   int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
    // 名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 楼层名称
    FloorName   string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
    // id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 高度
    Height   string `json:"height,omitempty" xml:"height,omitempty"`
}
