package paimai

// ItemMateriaValueDO 
type ItemMateriaValueDO struct {

    // 材质值名称
    Name   string `json:"name,omitempty"`

    // 当前材质值，是否需要填写含量值。比如：棉 是需要填写含量值，而牛皮 是不需要填写含量值的
    NeedContentNumber   bool `json:"need_content_number,omitempty"`

}
