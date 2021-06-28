package hotel

// TabInfo 
type TabInfo struct {

    // 标签的态度 1好 0中性 -1 差
    
    Attitude   int64 `json:"attitude,omitempty" xml:"attitude,omitempty"`
    

    // tab是否点击
    
    IsClick   bool `json:"is_click,omitempty" xml:"is_click,omitempty"`
    

    // tab编码
    
    TabCode   string `json:"tab_code,omitempty" xml:"tab_code,omitempty"`
    

    // tab下的统计数据
    
    TabDetail   string `json:"tab_detail,omitempty" xml:"tab_detail,omitempty"`
    

    // tabId
    
    TabId   int64 `json:"tab_id,omitempty" xml:"tab_id,omitempty"`
    

    // tab名称
    
    TabName   string `json:"tab_name,omitempty" xml:"tab_name,omitempty"`
    

    // tab对应的埋点名
    
    TabTrack   string `json:"tab_track,omitempty" xml:"tab_track,omitempty"`
    

    // 标签的类型 0正常 1热词 2房型 3度假
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

}
