package viapi

// LogoData 
type LogoData struct {

    // 识别出的logo类型，取值为TV （台标）
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 识别出的logo名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 以图片左上角为坐标原点，logo区域左上角到y轴距离。
    
    X   int64 `json:"x,omitempty" xml:"x,omitempty"`
    

    // 以图片左上角为坐标原点，logo区域左上角到x轴距离。
    
    Y   int64 `json:"y,omitempty" xml:"y,omitempty"`
    

    // logo区域宽度
    
    Width   int64 `json:"width,omitempty" xml:"width,omitempty"`
    

    // logo区域高度
    
    Height   int64 `json:"height,omitempty" xml:"height,omitempty"`
    

}
