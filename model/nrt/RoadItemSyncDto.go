package nrt

// RoadItemSyncDto 
type RoadItemSyncDto struct {
    // 数据
    List   []RoadItemSaveDto `json:"list,omitempty" xml:"list>road_item_save_dto,omitempty"`
    // 设备Sn
    DeviceSn   string `json:"device_sn,omitempty" xml:"device_sn,omitempty"`
}
