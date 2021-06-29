package axintrade

// AxinPayCheckSignDTO 
type AxinPayCheckSignDTO struct {
    // 业务场景
    BizScene   string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
    // 验签数据
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
}
