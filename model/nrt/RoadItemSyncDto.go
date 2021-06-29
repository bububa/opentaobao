package nrt

// RoadItemSyncDto 
type RoadItemSyncDto struct {

    // 数据
    
    List   []RoadItemSaveDto `json:"list,omitempty" xml:"list,omitempty"`
    

    // 设备Sn
    
    DeviceSn   string `json:"device_sn,omitempty" xml:"device_sn,omitempty"`
    

}
